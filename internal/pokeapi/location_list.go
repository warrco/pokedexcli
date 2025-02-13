package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLocations(url string, config *Locations) error {
	if url == "" {
		url = baseURL + "/location-area"
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making GET request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("bad status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, config)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}
	return nil
}
