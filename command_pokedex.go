package main

import (
	"fmt"
)

func commandPokedex(cfg *config, name string) error {
	fmt.Println("Your Pokedex:")
	for _, val := range cfg.pokeDex {
		fmt.Println(" - ", val.Name)
	}
	return nil
}
