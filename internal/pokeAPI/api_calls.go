package pokeAPI

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

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

type Client struct {
    httpClient http.Client
}


func NewClient(timeout time.Duration) Client {
    return Client{
        httpClient: http.Client{
            Timeout: timeout,
        },
    }
}

func getRequest(url string) (response []byte, err error) {
    // prepare the request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {return}

    // create a client and make the request
    client := http.Client{}
    resp, err := client.Do(req)
    if err != nil {return}
    defer resp.Body.Close()

    // convert the response
    response, err = io.ReadAll(resp.Body)

    return
}

func GetLocations(url string) (loc locationResponse, err error) {
    dataset, err := getRequest(url)
    if err != nil {return}

    err = json.Unmarshal(dataset, &loc)
    return
}
