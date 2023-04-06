package main

import (
	"flag"
	"time"

	"github.com/6rian/pokedexcli/cache"
	"github.com/6rian/pokedexcli/config"
	"github.com/6rian/pokedexcli/pokeapi"
)

func main() {
	var debugFlag bool
	flag.BoolVar(&debugFlag, "debug", false, "Enables debugging commands")
	flag.Parse()

	cfg := config.Config{
		PokeApiClient: pokeapi.New(),
		Cache:         cache.NewCache(3 * time.Minute),
		DebugMode:     debugFlag,
	}
	StartRepl(&cfg)
}
