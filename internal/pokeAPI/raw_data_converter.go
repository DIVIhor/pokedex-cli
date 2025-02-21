package pokeAPI

import "encoding/json"

// map locations responce structure
type locationResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}


func ReadJson(rawData []byte) (dataset locationResponse, err error) {
	err = json.Unmarshal(rawData, &dataset)
	return
}