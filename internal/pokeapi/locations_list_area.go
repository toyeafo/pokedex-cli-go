package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationsArea(name string) (locationAreaExplore, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		locationResp := locationAreaExplore{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return locationAreaExplore{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaExplore{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaExplore{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationAreaExplore{}, err
	}

	locationResp := locationAreaExplore{}
	jsonerr := json.Unmarshal(body, &locationResp)
	if jsonerr != nil {
		return locationAreaExplore{}, jsonerr
	}

	c.cache.Add(url, body)
	return locationResp, nil
}
