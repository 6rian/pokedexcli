package pokeapi

import (
	"errors"
	"io"
	"net/http"
	"time"
)

func New() *PokeApiClient {
	return &PokeApiClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: "https://pokeapi.co/api/v2/",
	}
}

func (p *PokeApiClient) Fetch(url string) ([]byte, error) {
	if url == "" {
		return []byte{}, errors.New("missing URL to fetch")
	}

	resp, err := p.client.Get(url)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return []byte{}, errors.New("invalid location area")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (p *PokeApiClient) GetLocationAreasUrl() string {
	return p.baseUrl + "location-area/"
}

func (p *PokeApiClient) GetLocationAreaUrl(name string) string {
	return p.GetLocationAreasUrl() + name
}

func (p *PokeApiClient) GetPokemonUrl(name string) string {
	return p.baseUrl + "pokemon/" + name
}
