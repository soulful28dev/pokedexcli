package internal

import (
	"testing"
	"time"
)

func TestAddCache(t *testing.T) {
	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
			inputVal: []byte("val test"),
		},
	}

	var timeDuration time.Duration = 5 * time.Minute
	var cache = NewCache(timeDuration)

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputVal)
		_, ok := cache.Get(c.inputKey)
		if !ok {
			t.Errorf("AddCache %s - failed", c.inputKey)
		}
	}
}
