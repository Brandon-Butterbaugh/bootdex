package main

import (
	"time"

	"github.com/Brandon-Butterbaugh/bootdex/internal/pokeapi"
	"github.com/Brandon-Butterbaugh/bootdex/internal/pokecache"
)

func main() {
	pokeCache := pokecache.NewCache(5 * time.Second)
	pokeClient := pokeapi.NewClient(5*time.Second, pokeCache)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
