package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	// Check Arguments
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	// Get Pokemon Info from Name argument
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	// Catch the Pokemon!
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	random := rand.IntN(pokemon.BaseExperience)
	if random > 40 {
		fmt.Printf("%s got away...\n", pokemon.Name)
		return nil
	}

	fmt.Printf("You caught %s!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}
