package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListPokemon -
func (c *Client) ListPokemon(location string) (RespShallowPokemon, error) {
	url := baseURL + "/location-area/" + location

	if data, ok := c.pokeCache.Get(url); ok {
		pokemonResp := RespShallowPokemon{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return RespShallowPokemon{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RespShallowPokemon{}, fmt.Errorf("status %d", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	c.pokeCache.Add(url, dat)

	pokemonResp := RespShallowPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	return pokemonResp, nil
}
