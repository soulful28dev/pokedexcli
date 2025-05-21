package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type PokemonInfoStat struct {
	Stat struct {
		Name string `json:"name"`
	} `json:"stat"`
	Effort   int `json:"effort"`
	BaseStat int `json:"base_stat"`
}

type PokemonInfoType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type PokemonInfo struct {
	Id             int               `json:"id"`
	Name           string            `json:"name"`
	BaseExperience int               `json:"base_experience"`
	Height         int               `json:"height"`
	Weight         int               `json:"weight"`
	Stats          []PokemonInfoStat `json:"stats"`
	Types          []PokemonInfoType `json:"types"`
}

func commandCatch(args ...string) error {
	if len(args) <= 0 {
		return errors.New("errors on command catch args")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	resp, resErr := http.Get(url)
	if resErr != nil {
		return resErr
	}
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data PokemonInfo
	if err := json.Unmarshal(respData, &data); err != nil {
		return err
	}

	rand := rand.Intn(2)
	if rand == 0 {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		userPokedex[pokemonName] = data
		fmt.Printf("%s was caught!\n", pokemonName)
	}

	return nil
}
