package main

import (
	"fmt"

	"github.com/kitaclysm/gator/internal/config"
	"github.com/kitaclysm/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	names map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if cmdval, ok := c.names[cmd.Name]; ok {
		return cmdval(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.Name)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.names[name] = f
}