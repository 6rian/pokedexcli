package main

import (
	"time"

	"github.com/6rian/pokedexcli/cache"
	"github.com/6rian/pokedexcli/config"
	"github.com/6rian/pokedexcli/pokeapi"
)

func main() {
	cfg := config.Config{
		PokeApiClient: pokeapi.New(),
		Cache:         cache.NewCache(3 * time.Minute),
	}
	StartRepl(&cfg)
}
