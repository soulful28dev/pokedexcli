package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapBack(args ...string) error {
	url := cfg.Previous

	if url == "" {
		fmt.Println("No previous page available.")
		return nil
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
