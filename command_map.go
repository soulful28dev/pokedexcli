package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	Next     string
	Previous string
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(args ...string) error {
	url := cfg.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}

	var respData []byte

	cacheResp, ok := cache.Get(url)

	if ok {
		respData = cacheResp
	} else {
		resp, resErr := http.Get(url)
		if resErr != nil {
			return resErr
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		respData = body
		cache.Add(url, respData)
		defer resp.Body.Close()
	}

	var data LocationAreaResponse
	if err := json.Unmarshal(respData, &data); err != nil {
		return err
	}

	// Update config pagination URLs
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, loc := range data.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
