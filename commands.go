package main

import (
	"fmt"
	"github.com/6rian/pokedexcli/pokeapi"
	"os"
)

func CommandHelp() error {
	fmt.Printf("\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func CommandExit() error {
	defer os.Exit(0)
	fmt.Printf("\nByebye!\n")
	return nil
}

func CommandMap() error {
	pokeapi.Test()
	pokeapiClient := pokeapi.New()
	err := pokeapiClient.GetMap()
	if err != nil {
		return err
	}
	return nil
}
