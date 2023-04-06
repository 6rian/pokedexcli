package commands

import (
	"fmt"

	"github.com/6rian/pokedexcli/config"
)

func commandDumpCache(cfg *config.Config, args CliCommandArgs) error {
	fmt.Printf("[DEBUG] Dumping the cache...\n")

	for key := range cfg.Cache.Entries() {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
