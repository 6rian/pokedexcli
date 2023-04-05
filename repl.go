package main

import (
	"bufio"
	"fmt"
	"github.com/6rian/pokedexcli/commands"
	"log"
	"os"
)

func prompt() {
	fmt.Printf("Pokedex> ")
}

func StartRepl() {
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
			err := cmd.Callback()
			if err != nil {
				log.Fatalf("An error occurred: %v", err)
			}
		} else {
			fmt.Printf("Oops! Unknown command. Please try again or type 'help' for usage.\n")
		}
	}
}
