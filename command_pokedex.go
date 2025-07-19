package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	// Get Pokemon Info from Name argument in pokedex
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
