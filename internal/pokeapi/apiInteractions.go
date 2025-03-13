package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
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
func (c *Client) ResponseData(urlInput *string) ([]byte, error) {

	req, err := http.NewRequest("GET", *urlInput, nil)
	if err != nil {
		return []byte{}, fmt.Errorf("error making request")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("error fetching response")
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error parsing response")
	}

	return respData, nil
}

func UnmarshalSliceOfBytesLocationAreas(respData []byte) (locationAreaEndpointData, error) {
	locationAreas := locationAreaEndpointData{}

	err := json.Unmarshal(respData, &locationAreas)
	if err != nil {
		return locationAreaEndpointData{}, fmt.Errorf("error decoding response")
	}

	return locationAreas, nil
}

func UnmarshalSliceOfBytesLocationAreasDetails(respData []byte) (locationAreaDetails, error) {
	locationDetails := locationAreaDetails{}

	err := json.Unmarshal(respData, &locationDetails)
	if err != nil {
		return locationAreaDetails{}, fmt.Errorf("error decoding response")
	}

	return locationDetails, nil
}

func UnmarshalSliceOfBytesPokemonDetails(respData []byte) (PokemonDetails, error) {
	pokemonInfo := PokemonDetails{}

	err := json.Unmarshal(respData, &pokemonInfo)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("error decoding response")
	}

	return pokemonInfo, nil
}
