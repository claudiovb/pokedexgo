package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/claudiovb/pokedexcli/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	pokeapiClient  *api.Client
	nextApiUrl     string
	previousApiUrl string
	pokemonWallet  map[string]api.Pokemon
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	config := &Config{
		pokeapiClient:  api.NewClient(5*time.Second, 5*time.Minute),
		nextApiUrl:     "https://pokeapi.co/api/v2/location-area",
		previousApiUrl: "https://pokeapi.co/api/v2/location-area",
		pokemonWallet:  map[string]api.Pokemon{},
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		args := []string{}
		commandName := words[0]
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command. Type 'exit' to quit.")
			continue
		}
		err := command.callback(config, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the location areas in forward mode",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the location areas in backward mode",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore the pokemons on the location areas",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List the pokemons in the pokedex",
			callback:    commandPokedex,
		},
	}
}
