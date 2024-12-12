package api

import (
	"net/http"
	"time"

	"github.com/claudiovb/pokedexcli/internal/pokecache"
)

type Client struct {
	Cache      *pokecache.Cache
	HttpClient *http.Client
}

func NewClient(interval, cacheInterval time.Duration) *Client {
	return &Client{
		Cache:      pokecache.NewCache(cacheInterval),
		HttpClient: &http.Client{Timeout: interval},
	}
}
