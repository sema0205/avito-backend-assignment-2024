package client

import (
	"github.com/patrickmn/go-cache"
	memory "github.com/sema0205/avito-backend-assignment-2024/pkg/cache"
	"time"
)

type Client struct {
	cache *cache.Cache
}

func NewCacheClient(defaultTTl time.Duration, cleanupTTL time.Duration) memory.Provider {
	return &Client{
		cache: cache.New(defaultTTl, cleanupTTL),
	}
}

func (c *Client) Set(key string, value interface{}) {
	c.cache.Set(key, value, cache.DefaultExpiration)
}

func (c *Client) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

func (c *Client) Delete(key string) {
	c.cache.Delete(key)
}
