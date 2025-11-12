package pokeapi

import (
	"net/http"
	"time"

	"github.com/Brandon-Butterbaugh/bootdex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
	Pokemon    map[string]Pokemon
}

// NewClient -
func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: cache,
		Pokemon:   make(map[string]Pokemon),
	}
}
