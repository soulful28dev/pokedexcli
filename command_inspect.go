package main

import (
	"errors"
	"fmt"
)

func (pInfo PokemonInfo) displayInfo() {
	fmt.Printf("Name: %s\n", pInfo.Name)
	fmt.Printf("Height: %v\n", pInfo.Height)
	fmt.Printf("Weight: %v\n", pInfo.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pInfo.Stats {
		fmt.Printf("\t-%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")

	for _, infoType := range pInfo.Types {
		fmt.Printf("\t- %s\n", infoType.Type.Name)
	}
}

func commandInspect(args ...string) error {
	if len(args) <= 0 {
		return errors.New("errors on command catch args")
	}

	pokemonName := args[0]
	existPokemon, ok := userPokedex[pokemonName]
	if ok {
		existPokemon.displayInfo()
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
