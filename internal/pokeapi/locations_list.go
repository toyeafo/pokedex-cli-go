package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (locationArea, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationArea{}, err
	}

	locationResp := locationArea{}
	jsonerr := json.Unmarshal(body, &locationResp)
	if jsonerr != nil {
		return locationArea{}, jsonerr
	}

	return locationResp, nil
}
