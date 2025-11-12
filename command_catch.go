package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 2 {
		fmt.Println("Please specify a Pokemon to try catching!")
		fmt.Println("Example: catch <pokemon>")
		return nil
	}
	pokemonResp, err := cfg.pokeapiClient.PokemonCatch(args[1])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)
	throw := rand.IntN(350)
	if throw > pokemonResp.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		cfg.pokeapiClient.Pokemon[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}
	return nil
}
