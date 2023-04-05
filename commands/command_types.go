package commands

import "github.com/6rian/pokedexcli/config"

type CliCommand struct {
	name        string
	description string
	Callback    func(cfg *config.Config) error
}

type CliCommandMap map[string]CliCommand
