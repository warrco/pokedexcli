package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocationInfo(locationName string) (LocationInfo, error) {
	url := baseURL + "/location-area/" + locationName

	if data, ok := c.cache.Get(url); ok {
		pokemonList := LocationInfo{}
		err := json.Unmarshal(data, &pokemonList)
		if err != nil {
			return LocationInfo{}, fmt.Errorf("error unmarshaling JSON: %w", err)
		}
		return pokemonList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationInfo{}, fmt.Errorf("error making GET request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationInfo{}, fmt.Errorf("no data returned: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationInfo{}, fmt.Errorf("error reading response body: %w", err)
	}

	pokemonList := LocationInfo{}
	err = json.Unmarshal(body, &pokemonList)
	if err != nil {
		return LocationInfo{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	c.cache.Add(url, body)
	return pokemonList, nil
}
