package main

import (
	"fmt"

	"github.com/kitaclysm/gator/internal/config"
)

type state struct {
	db	*database.Queries
	cfg	*config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	names map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if cmdval, ok := c.names[cmd.name]; ok {
		return cmdval(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.name)
}

func (c *commands) register(name string, f func(*state, command) error ) {
	c.names[name] = f
}