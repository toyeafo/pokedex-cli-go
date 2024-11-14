package main

import (
	"time"

	"github.com/toyeafo/pokedex-cli-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 3*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startComm(cfg)
}
