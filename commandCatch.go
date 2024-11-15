package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {
	if name == "" {
		return errors.New("provide a pokemon name")
	}

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Println("Throwing a ball at " + name + "...")

	var chances int = rand.Int()

	if chances > pokemonResp.BaseExperience {
		fmt.Println(name + " was caught")
	} else {
		fmt.Println(name + " escaped")
	}

	cfg.pokeDex[pokemonResp.Name] = pokemonResp
	return nil
}
