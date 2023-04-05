package commands

type CliCommand struct {
	name        string
	description string
	Callback    func() error
}

type CliCommandMap map[string]CliCommand
