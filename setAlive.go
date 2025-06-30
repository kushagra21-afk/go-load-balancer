package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Container struct {
	mu           sync.RWMutex
	parsed_url   *url.URL
	ReverseProxy *httputil.ReverseProxy
	alive        bool
}

func newContainer(URL string) (*Container, error) {
	parsed_url, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	return &Container{
		parsed_url:   parsed_url,
		ReverseProxy: httputil.NewSingleHostReverseProxy(parsed_url),
		alive:        true,
	}, nil
}
func (c *Container) isAlive() bool {
	lock := c.mu.RLocker()
	lock.Lock()
	defer lock.Unlock()
	return c.alive
}
func (c *Container) setAlive(status bool) {
	lock := c.mu.RLocker()
	lock.Lock()
	defer lock.Unlock()
	c.alive = status
}
