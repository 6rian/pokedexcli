package pokeapi

import (
	"io"
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

func (p *PokeApiClient) GetDefaultLocationAreasUrl() string {
	return p.baseUrl + "location-area/"
}

func (p *PokeApiClient) FetchLocationAreas(url string) ([]byte, error) {
	if url == "" {
		url = p.GetDefaultLocationAreasUrl()
	}

	resp, err := p.client.Get(url)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	// TODO check the response status code

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
