package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func ResponseData(url string) (locationAreaEndpointData, error) {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		return locationAreaEndpointData{}, fmt.Errorf("error sending request : %w", err)
// 	}

// 	defer res.Body.Close()

// 	configData := locationAreaEndpointData{}

// 	decoder := json.NewDecoder(res.Body)
// 	err = decoder.Decode(&configData)
// 	if err != nil {
// 		return locationAreaEndpointData{}, fmt.Errorf("error decoding response : %w", err)
// 	}

// 	return configData, nil
// }

// Sends GET request and decodes the incoming data into a slice of "Config" struct
func (c *Client) ResponseData(urlInput *string) (locationAreaEndpointData, error) {

	req, err := http.NewRequest("GET", *urlInput, nil)
	if err != nil {
		return locationAreaEndpointData{}, fmt.Errorf("error making request")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaEndpointData{}, fmt.Errorf("error fetching response")
	}

	defer resp.Body.Close()

	locationAreas := locationAreaEndpointData{}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locationAreas)
	if err != nil {
		return locationAreaEndpointData{}, fmt.Errorf("error decoding response")
	}

	return locationAreas, nil
}
