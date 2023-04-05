package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PokeApi struct {
	client *http.Client
}

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func New() *PokeApi {
	return &PokeApi{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (p *PokeApi) GetMap() error {
	var url = "https://pokeapi.co/api/v2/location-area/"

	resp, err := p.client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var target interface{}
	json.NewDecoder(resp.Body).Decode(&target)
	fmt.Printf("Results: \n%v\n", target)
	return nil
}

func Test() {
	fmt.Printf("\nrunning pokeapi.Test\n")
}
