package pokeAPI

import (
	"io"
	"net/http"
	"time"

	"github.com/DIVIgor/pokedex-cli/internal/pokecache"
)


type Client struct {
    cache pokecache.Cache
    httpClient http.Client
}


func NewClient(timeout, cachingDuration time.Duration) Client {
    return Client{
        cache: *pokecache.NewCache(cachingDuration),
        httpClient: http.Client{
            Timeout: timeout,
        },
    }
}

func (c *Client) getRequest(url string) (response []byte, err error) {
    // prepare the request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {return}

    resp, err := c.httpClient.Do(req)
    if err != nil {return}
    defer resp.Body.Close()

    // convert the response
    response, err = io.ReadAll(resp.Body)

    return
}

func (c *Client) GetLocations(url string) (dataset []byte, err error) {
    dataset, cached := c.cache.Get(url)
    if !cached {
        dataset, err = c.getRequest(url)
        if err != nil {
            return
        }
        c.cache.Add(url, dataset)
    }

    return
}
