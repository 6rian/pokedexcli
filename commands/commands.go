package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

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
	fmt.Printf("\nUsage:\n\n")
	for _, cmd := range GetCommands(cfg.DebugMode) {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
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

	var locations pokeapi.LocationAreasResp
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
		return errors.New("missing a valid location area")
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

	var area pokeapi.LocationAreaResp
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
