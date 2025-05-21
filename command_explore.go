package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationAreaPokemonResponse struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon Pokemon
	} `json:"pokemon_encounters"`
}

func commandExplore(args ...string) error {
	if len(args) <= 0 {
		return errors.New("errors on command explore args")
	}

	locationArea := args[0]

	url := "https://pokeapi.co/api/v2/location-area/" + locationArea

	resp, resErr := http.Get(url)
	if resErr != nil {
		return resErr
	}
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data LocationAreaPokemonResponse
	if err := json.Unmarshal(respData, &data); err != nil {
		return err
	}

	for _, pokemon := range data.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
