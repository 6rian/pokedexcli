package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/6rian/pokedexcli/commands"
	"github.com/6rian/pokedexcli/config"
)

func prompt() {
	fmt.Printf("Pokedex> ")
}

func StartRepl(cfg *config.Config) {
	withDebugging := cfg.DebugMode
	cmds := commands.GetCommands(withDebugging)
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
				fmt.Printf("[!] ERROR: %v\n", err)
			}
		} else {
			fmt.Printf("Oops! Unknown command. Please try again or type 'help' for usage.\n")
		}
	}
}
