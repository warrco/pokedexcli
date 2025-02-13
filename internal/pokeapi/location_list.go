package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchLocations(pageURL *string) (AreaLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, ok := c.cache.Get(url); ok {
		areaLocations := AreaLocations{}
		err := json.Unmarshal(data, &areaLocations)
		if err != nil {
			return AreaLocations{}, fmt.Errorf("error unmarshaling JSON: %w", err)
		}
		return areaLocations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return AreaLocations{}, fmt.Errorf("error making GET request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AreaLocations{}, fmt.Errorf("no data returned: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaLocations{}, fmt.Errorf("error reading response body: %w", err)
	}

	areaLocations := AreaLocations{}
	err = json.Unmarshal(body, &areaLocations)
	if err != nil {
		return AreaLocations{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	c.cache.Add(url, body)
	return areaLocations, nil
}
