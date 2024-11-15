package main

import "fmt"

func commandHelp(cfg *config, area string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()
	for _, cmd := range commandInp() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
