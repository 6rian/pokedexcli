package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sort"

	"github.com/6rian/pokedexcli/config"
	"github.com/6rian/pokedexcli/pokeapi"
)

func GetCommands(withDebugging bool) CliCommandMap {
	cm := CliCommandMap{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			Callback:    commandMapb,
		},
		"explore": {
			name:        "explore [area]",
			description: "Explore Pokemon available in the area",
			Callback:    commandExplore,
		},
		"catch": {
			name:        "catch [pokemon]",
			description: "Try to catch the Pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect [pokemon]",
			description: "View all details about the Pokemon",
			Callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all the Pokemon in your pokedex",
			Callback:    commandPokedex,
		},
	}

	if withDebugging {
		cm["debug.dumpCache"] = CliCommand{
			name:        "debug.dumpCache",
			description: "Dumps all entries in the cache",
			Callback:    commandDumpCache,
		}
	}

	return cm
}

func commandHelp(cfg *config.Config, args CliCommandArgs) error {
	commands := GetCommands(cfg.DebugMode)
	var names []string
	for k := range commands {
		names = append(names, k)
	}
	sort.Strings(names)

	fmt.Printf("\nUsage:\n\n")
	for _, v := range names {
		fmt.Printf(" - %s: %s\n", commands[v].name, commands[v].description)
	}
	return nil
}

func commandExit(cfg *config.Config, args CliCommandArgs) error {
	defer os.Exit(0)
	fmt.Printf("Byebye!\n")
	return nil
}

func commandMap(cfg *config.Config, args CliCommandArgs) error {
	var url string
	if cfg.Next == nil {
		url = cfg.PokeApiClient.GetLocationAreasUrl()
	} else {
		url = *cfg.Next
	}
	return getMap(url, cfg)
}

func commandMapb(cfg *config.Config, args CliCommandArgs) error {
	var url string
	if cfg.Prev == nil {
		return errors.New("you're already at the beginning, you can't go back")
	} else {
		url = *cfg.Prev
	}
	return getMap(url, cfg)
}

func getMap(url string, cfg *config.Config) error {
	var resp []byte
	resp, exists := cfg.Cache.Get(url)
	if !exists {
		var err error
		resp, err = cfg.PokeApiClient.Fetch(url)
		if err != nil {
			return err
		}
	}

	var locations pokeapi.LocationAreas
	err := json.Unmarshal(resp, &locations)
	if err != nil {
		return err
	}

	if !exists {
		cfg.Cache.Add(url, resp)
	}

	cfg.Next = locations.Next
	cfg.Prev = locations.Previous

	for _, area := range locations.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}

func commandExplore(cfg *config.Config, args CliCommandArgs) error {
	if len(args) != 1 {
		return errors.New("please enter a single location area to explore")
	}

	url := cfg.PokeApiClient.GetLocationAreaUrl(args[0])

	var resp []byte
	resp, exists := cfg.Cache.Get(url)
	if !exists {
		var err error
		resp, err = cfg.PokeApiClient.Fetch(url)
		if err != nil {
			return err
		}
	}

	var area pokeapi.LocationArea
	err := json.Unmarshal(resp, &area)
	if err != nil {
		return err
	}

	if !exists {
		cfg.Cache.Add(url, resp)
	}

	if len(area.PokemonEncounters) == 0 {
		fmt.Printf("No Pokenmon found.\n")
	} else {
		fmt.Printf("Found Pokemon:\n")
		for _, encounters := range area.PokemonEncounters {
			fmt.Printf(" - %s\n", encounters.Pokemon.Name)
		}
	}

	return nil
}

func commandCatch(cfg *config.Config, args CliCommandArgs) error {
	if len(args) != 1 {
		return errors.New("please enter a single pokemon to catch")
	}

	name := args[0]
	url := cfg.PokeApiClient.GetPokemonUrl(name)

	if _, caught := cfg.Pokedex[name]; caught {
		fmt.Printf("%s was already caught\n", name)
		return nil
	} else {
		fmt.Printf("Throwing a Pokeball at %s...\n", name)
	}

	var resp []byte
	resp, exists := cfg.Cache.Get(url)
	if !exists {
		var err error
		resp, err = cfg.PokeApiClient.Fetch(url)
		if err != nil {
			return err
		}
	}

	var pokemon pokeapi.Pokemon
	err := json.Unmarshal(resp, &pokemon)
	if err != nil {
		return err
	}

	if !exists {
		cfg.Cache.Add(url, resp)
	}

	caught := catch(pokemon.BaseExperience)
	if caught {
		cfg.Pokedex[name] = pokemon
		fmt.Printf("%s was caught!\n", name)
		fmt.Printf("You may now inspect it with the inspect command.\n")
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func catch(experience int) bool {
	// Max base experience is 608
	chance := rand.Intn(750)
	return chance >= experience
}

func commandInspect(cfg *config.Config, args CliCommandArgs) error {
	if len(args) != 1 {
		return errors.New("please enter a single pokemon to inspect")
	}

	name := args[0]

	p, caught := cfg.Pokedex[name]
	if !caught {
		fmt.Printf("You have not caught that pokemon!\n")
		return nil
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Printf("Stats:\n")
	for _, s := range p.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, t := range p.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config.Config, args CliCommandArgs) error {
	if len(cfg.Pokedex) == 0 {
		fmt.Printf("Your Pokedex is empty. Go catch some Pokemon!\n")
		return nil
	}

	fmt.Printf("Your Pokedex:\n")
	for k := range cfg.Pokedex {
		fmt.Printf(" - %s\n", k)
	}

	return nil
}
