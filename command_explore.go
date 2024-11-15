package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, area string) error {
	if area == "" {
		return errors.New("a location name should be provided")
	}

	locationResp, err := cfg.pokeapiClient.ListLocationsArea(area)
	if err != nil {
		return err
	}

	for _, val := range locationResp.PokemonEncounters {
		fmt.Println(val.Pokemon.Name)
	}

	return nil
}
