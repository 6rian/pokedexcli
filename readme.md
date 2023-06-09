# PokedexCLI

PokedexCLI is a command line REPL for interacting with [PokeApi](https://pokeapi.co/). It includes commands for navigating through locations in the Pokemon world, exploring those locations for Pokemon, and catching and inspecting Pokemon in your Pokedex.

To get started, type `help` to see a list of available commands.

## Requirements
- Go version 1.18 or higher

## Installation

```bash
# Install
go install github.com/6rian/pokedexcli@latest

# Run application (Assumes $GOPATH is set)
pokedexcli
```

## Build From Source

```bash
# Clone the repository
git clone git@github.com:6rian/pokedexcli.git
cd pokedexcli

# Build
go build

# Run application
./pokedexcli

# Run tests
go test ./...
```
