package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		locationResp := PokemonInfo{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return PokemonInfo{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	locationResp := PokemonInfo{}
	jsonerr := json.Unmarshal(body, &locationResp)
	if jsonerr != nil {
		return PokemonInfo{}, jsonerr
	}

	c.cache.Add(url, body)
	return PokemonInfo{}, nil
}
