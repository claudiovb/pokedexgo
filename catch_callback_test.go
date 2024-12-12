package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/claudiovb/pokedexcli/internal/api"
	"github.com/claudiovb/pokedexcli/internal/pokecache"
)

type MockPokeapiClient struct{}

func (m *MockPokeapiClient) GetPokemon(name string) (api.Pokemon, error) {
	return api.Pokemon{BaseExperience: 100}, nil
}

func createMockClient(mockServer *httptest.Server) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			DialTLS: func(network, addr string) (net.Conn, error) {
				// Always redirect to mock server's URL
				mockURL, _ := url.Parse(mockServer.URL)
				return net.Dial("tcp", mockURL.Host)
			},
		},
	}
}

func TestCatchPokemon(t *testing.T) {
	fmt.Println("TestCatchPokemon")
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Calling on mock Server")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"base_experience": 1}`))
	}))
	defer mockServer.Close()
	pokemonName := "pikachu"
	mockClient := createMockClient(mockServer)
	pokemon := api.Pokemon{BaseExperience: 1}
	client := &api.Client{
		Cache:      pokecache.NewCache(5 * time.Second),
		HttpClient: mockClient,
	}
	config := &Config{
		pokeapiClient: client,
		pokemonWallet: map[string]api.Pokemon{},
	}

	err := commandCatch(config, pokemonName)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if config.pokemonWallet[pokemonName].BaseExperience != pokemon.BaseExperience {
		t.Errorf("Expected base experience %v, got %v", pokemon.BaseExperience, config.pokemonWallet[pokemonName].BaseExperience)
	}
}
