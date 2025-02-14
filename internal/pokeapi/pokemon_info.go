package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if data, ok := c.cache.Get(url); ok {
		pokemonList := Pokemon{}
		err := json.Unmarshal(data, &pokemonList)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshaling JSON: %w", err)
		}
		return pokemonList, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error making GET request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("no data returned: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response body: %w", err)
	}

	pokemonList := Pokemon{}
	err = json.Unmarshal(body, &pokemonList)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	c.cache.Add(url, body)
	return pokemonList, nil
}
