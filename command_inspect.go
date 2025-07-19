package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	// Check Arguments
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	// Get Pokemon Info from Name argument in pokedex
	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// Print Pokemon details
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil
}
