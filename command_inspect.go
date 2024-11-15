package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	key, prs := cfg.pokeDex[name]
	if !prs {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name: ", key.Name)
	fmt.Println("Height: ", key.Height)
	fmt.Println("Weight: ", key.Weight)
	fmt.Println("Stat:")
	for _, stat := range key.Stats {
		fmt.Println("-", stat.Stat.Name, ": ", stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range key.Types {
		fmt.Println("- ", val.Type.Name)
	}
	return nil
}
