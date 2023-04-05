package commands

import (
	"fmt"
	"github.com/6rian/pokedexcli/pokeapi"
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

func commandHelp() error {
	fmt.Printf("\nUsage:\n\n")
	for _, cmd := range GetCommands() {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}

func commandExit() error {
	defer os.Exit(0)
	fmt.Printf("Byebye!\n")
	return nil
}

func commandMap() error {
	pokeapi.Test()
	pokeapiClient := pokeapi.New()
	err := pokeapiClient.GetMap()
	if err != nil {
		return err
	}
	return nil
}
