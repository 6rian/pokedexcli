package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/6rian/pokedexcli/commands"
	"github.com/6rian/pokedexcli/config"
)

func StartRepl(cfg *config.Config) {
	withDebugging := cfg.DebugMode
	cmds := commands.GetCommands(withDebugging)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	for {
		prompt()

		// Read input
		scanner.Scan()
		input := parseInput(scanner.Text())

		// Evaluate and execute command
		var cmdName string = input[0]
		var cmdArgs commands.CliCommandArgs = input[1:]
		if cmd, ok := cmds[cmdName]; ok {
			err := cmd.Callback(cfg, cmdArgs)
			if err != nil {
				printError(err)
			}
		} else {
			printError(errors.New("Unknown command. Please try again or type 'help' for usage."))
		}
	}
}

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	slice := strings.Split(input, " ")
	return slice
}

func printError(err error) {
	fmt.Printf("[!] ERROR: %s\n", err.Error())
}

func prompt() {
	fmt.Printf("\nPokedex> ")
}
