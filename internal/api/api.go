package api

import (
	"encoding/json"
	"io"
)

func (c *Client) GetLocationsArea(url string) (LocationArea, error) {
	if val, ok := c.Cache.Get(url); ok {
		var locationArea LocationArea
		err := json.Unmarshal(val, &locationArea)
		if err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}
	res, err := c.HttpClient.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()
	var locationArea LocationArea
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	c.Cache.Add(url, body)
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

func (c *Client) GetPokemons(url string) (Pokemons, error) {
	if val, ok := c.Cache.Get(url); ok {
		var pokemons Pokemons
		err := json.Unmarshal(val, &pokemons)
		if err != nil {
			return Pokemons{}, err
		}
		return pokemons, nil
	}
	res, err := c.HttpClient.Get(url)
	if err != nil {
		return Pokemons{}, err
	}
	defer res.Body.Close()
	var pokemons Pokemons
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemons{}, err
	}
	c.Cache.Add(url, body)
	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		return Pokemons{}, err
	}
	return pokemons, nil

}

func (c *Client) GetPokemon(url string) (Pokemon, error) {
	if val, ok := c.Cache.Get(url); ok {
		var pokemon Pokemon
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	res, err := c.HttpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	// fmt.Println("Response status:", req.Status)
	var pokemon Pokemon
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.Cache.Add(url, body)
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
