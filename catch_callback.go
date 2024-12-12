package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a pokemon to catch")
		return nil
	}
	fmt.Println("Throwing a Pokeball at", args[0]+"...")
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", args[0])
	pokemon, err := config.pokeapiClient.GetPokemon(url)
	if err != nil {
		return err
	}
	baseExperience := pokemon.BaseExperience
	winProbability := calculateWinProbablity((baseExperience))

	luckWin := rand.Intn(baseExperience)
	if rand.Float64() < winProbability || luckWin == 1 {
		fmt.Println(args[0], "was caught!")
		config.pokemonWallet[args[0]] = pokemon
	} else {
		fmt.Println(args[0], "escaped!")
	}
	return nil
}

func calculateWinProbablity(prob int) float64 {
	return 0.5 + 0.5/(1+float64(prob)/100.0)

}
