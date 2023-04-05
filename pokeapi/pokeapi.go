package pokeapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type PokeApiClient struct {
	client  *http.Client
	baseUrl string
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

func New() *PokeApiClient {
	return &PokeApiClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: "https://pokeapi.co/api/v2/",
	}
}

func (p *PokeApiClient) FetchMap(url *string) (LocationAreasResp, error) {
	endpoint := "location-area/"

	if url == nil {
		tempUrl := p.baseUrl + endpoint
		url = &tempUrl
	}

	resp, err := p.client.Get(*url)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer resp.Body.Close()

	var r LocationAreasResp
	json.NewDecoder(resp.Body).Decode(&r)
	return r, nil
}
