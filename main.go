package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func newCliCommand(name, desc string, cb func() error) *cliCommand {
	return &cliCommand{
		name:        name,
		description: desc,
		callback:    cb,
	}
}

var commands = make(map[string]cliCommand)

func commandHelp() error {
	fmt.Printf("\nUsage:\n\n")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandExit() error {
	defer os.Exit(0)
	fmt.Printf("\nByebye!\n")
	return nil
}

func main() {
	// Register commands.
	commands["help"] = *newCliCommand("help", "Displays a help message", commandHelp)
	commands["exit"] = *newCliCommand("exit", "Exits the pokedex", commandExit)

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	// REPL.
	for {
		// Prompt.
		fmt.Printf("\nPokedex> ")

		// Read input.
		scanner.Scan()
		in := scanner.Text()

		// Evaluate and execute command.
		if cmd, ok := commands[in]; ok {
			err := cmd.callback()
			if err != nil {
				log.Fatalf("An error occurred: %v", err)
			}
		} else {
			fmt.Printf("\nOops! Unknown command. Please try again or type 'help' for usage.")
		}
	}
}
