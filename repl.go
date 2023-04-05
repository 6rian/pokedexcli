package main

import (
	"bufio"
	"fmt"
	"github.com/6rian/pokedexcli/commands"
	"github.com/6rian/pokedexcli/config"
	"os"
)

func prompt() {
	fmt.Printf("Pokedex> ")
}

func StartRepl(cfg *config.Config) {
	cmds := commands.GetCommands()
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	for {
		prompt()

		// Read input.
		scanner.Scan()
		input := scanner.Text()

		// Evaluate and execute command.
		if cmd, ok := cmds[input]; ok {
			err := cmd.Callback(cfg)
			if err != nil {
				fmt.Printf("An error occurred: %v", err)
			}
		} else {
			fmt.Printf("Oops! Unknown command. Please try again or type 'help' for usage.\n")
		}
	}
}
