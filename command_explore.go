package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) < 2 {
		fmt.Println("Please specify a location area to explore")
		fmt.Println("Example: explore <area_name>")
		return nil
	}

	pokemonResp, err := cfg.pokeapiClient.ListPokemon(args[1])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[1])
	fmt.Println("Found Pokemon:")
	for _, pok := range pokemonResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pok.Pokemon.Name)
	}

	return nil
}
