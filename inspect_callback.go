package main

import "fmt"

func commandInspect(cfg *Config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("Please provide a Pokemon to inspect")
		return nil
	}
	pokemon, ok := cfg.pokemonWallet[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Println("-"+stat.Stat.Name+":", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Println("-", pokemonType.Type.Name)
	}
	return nil
}
