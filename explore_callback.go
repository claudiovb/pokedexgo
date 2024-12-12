package main

import "fmt"

func commandExplore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide an area to explore")
		return nil
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", args[0])
	locationData, err := cfg.pokeapiClient.GetPokemons(url)
	if err != nil {
		fmt.Println("Error exploring area:", err)
		return err
	}
	fmt.Println("\n Exploring: ...", args[0])
	if len(locationData.PokemonEncounters) > 0 {
		for _, pokemonsData := range locationData.PokemonEncounters {
			fmt.Println(pokemonsData.Pokemon.Name)
		}
	} else {
		fmt.Println("No pokemons found in this area")
	}
	return nil
}
