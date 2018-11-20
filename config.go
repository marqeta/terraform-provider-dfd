package main

import (
	"github.com/marqeta/go-dfd/dfd"
)

type Config struct {
	DotPath string
}

func (c *Config) Client() (interface{}, error) {
	client := dfd.NewClient(c.DotPath)
	return client, nil
}
