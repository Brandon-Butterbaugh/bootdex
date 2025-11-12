package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for _, pok := range cfg.pokeapiClient.Pokemon {
		fmt.Printf(" - %s\n", pok.Name)
	}
	return nil
}
