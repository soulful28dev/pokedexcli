package main

import "fmt"

func commandPokedex(args ...string) error {
	fmt.Println("Your Pokedex:")
	for name := range userPokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
