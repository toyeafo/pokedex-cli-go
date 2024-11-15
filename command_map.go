package main

import (
	"fmt"
)

func commandMapf(cfg *config, area string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.previousLocationURL = locationResp.Previous

	for _, val := range locationResp.Results {
		fmt.Println(val.Name)
	}

	return nil
}

func commandMapb(cfg *config, area string) error {
	if cfg.previousLocationURL == nil {
		return fmt.Errorf("you are on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.previousLocationURL = locationResp.Previous

	for _, val := range locationResp.Results {
		fmt.Println(val.Name)
	}

	return nil
}
