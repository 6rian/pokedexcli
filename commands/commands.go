package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/6rian/pokedexcli/config"
	"github.com/6rian/pokedexcli/pokeapi"
)

func GetCommands() CliCommandMap {
	return CliCommandMap{
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
	}
}

func commandHelp(cfg *config.Config) error {
	fmt.Printf("\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func commandExit(cfg *config.Config) error {
	defer os.Exit(0)
	fmt.Printf("Byebye!\n")
	return nil
}

func commandMap(cfg *config.Config) error {
	var url string
	if cfg.Next == nil {
		url = cfg.PokeApiClient.GetDefaultLocationAreasUrl()
	} else {
		url = *cfg.Next
	}
	return getMap(url, cfg)
}

func commandMapb(cfg *config.Config) error {
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
		resp, err = cfg.PokeApiClient.FetchLocationAreas(url)
		if err != nil {
			return err
		}
	}

	var locations pokeapi.LocationAreasResp
	err := json.Unmarshal(resp, &locations)
	if err != nil {
		return err
	}

	cfg.Cache.Add(url, resp)
	cfg.Next = locations.Next
	cfg.Prev = locations.Previous

	for _, area := range locations.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}
