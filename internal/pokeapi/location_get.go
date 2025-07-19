package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	// Check the cache before HTTP request
	if val, ok := c.cache.Get(url); ok {
		locationsResp := Location{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return Location{}, err
		}
		return locationsResp, nil
	}

	// If cache not found, make new HTTP GET Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	// Do the request and retrieve the response
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	// Read the response body into dat
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	// Unmarshal the response data
	locationsResp := Location{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	// Add url to the cache and return struct, no errors
	c.cache.Add(url, dat)
	return locationsResp, nil
}
