package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Brandon-Butterbaugh/bootdex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println("Closing the Pokedex... Goodbye!")
			return
		}
		text := cleanInput(scanner.Text())
		if len(text) == 0 {
			continue
		}
		cmd, ok := getCommands()[text[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func cleanInput(text string) []string {
	temp := strings.TrimSpace(text)
	lower := strings.ToLower(temp)
	split := strings.Fields(lower)
	return split
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations in the world of Pokemon",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the world of Pokemon",
			callback:    commandMapb,
		},
	}
}
