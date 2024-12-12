package main

import "fmt"

func commandPokedex(cfg *Config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokemonWallet {
		fmt.Println("- ", pokemon.Name)
	}
	return nil
}
