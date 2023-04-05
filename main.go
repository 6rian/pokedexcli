package main

import (
	"github.com/6rian/pokedexcli/config"
	"github.com/6rian/pokedexcli/pokeapi"
)

func main() {
	cfg := config.Config{
		PokeApiClient: pokeapi.New(),
	}
	StartRepl(&cfg)
}
