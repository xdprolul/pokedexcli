package main

import (
	"fmt"
	"errors"
)

func commandPokedex(cfg *config,args ...string) error {
	if len(cfg.caughtPokemon)==0 {
		return errors.New("Your pokedex is empty!")
	} else {
		fmt.Println("Your Pokedex:")
	}

	for _,pokemon:=range cfg.caughtPokemon {
		fmt.Printf("  -%s\n",pokemon.Name)
	}
	return nil
}
