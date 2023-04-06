package config

import (
	"github.com/6rian/pokedexcli/cache"
	"github.com/6rian/pokedexcli/pokeapi"
)

type Config struct {
	Next          *string
	Prev          *string
	PokeApiClient *pokeapi.PokeApiClient
	Cache         cache.Cache
	DebugMode     bool
}
