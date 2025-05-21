package main

import (
	"time"

	"github.com/soulful28dev/pokedexcli/internal"
)

var timeDuration time.Duration = 5 * time.Minute
var cache = internal.NewCache(timeDuration)
var userPokedex = map[string]PokemonInfo{}

func main() {
	startRepl()
}
