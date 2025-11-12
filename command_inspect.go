package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 2 {
		fmt.Println("Please specify a Pokemon to inspect")
		fmt.Println("Example: inspect <pokemon>")
		return nil
	}

	for _, pok := range cfg.pokeapiClient.Pokemon {
		if pok.Name == args[1] {
			fmt.Printf("Height: %d\n", pok.Height)
			fmt.Printf("Weight: %d\n", pok.Weight)
			fmt.Println("Stats:")
			for _, stat := range pok.Stats {
				fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Println("Types:")
			for _, t := range pok.Types {
				fmt.Printf(" - %s\n", t.Type.Name)
			}
			return nil
		}
	}

	fmt.Printf("you have not caught a %s yet!", args[1])

	return nil
}
