package commands

import (
	"fmt"
	"github.com/6rian/pokedexcli/config"
	"os"
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
	results, err := cfg.PokeApiClient.FetchMap(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = results.Next
	cfg.Prev = results.Previous

	for _, area := range results.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}