package main

import (
	"fmt"

	"github.com/kitaclysm/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	names map[string]func(*state, command) error
}

// func handlerLogin(s *state, cmd command) error {
// 	if len(cmd.args) == 0 {
// 		return fmt.Errorf("no username provided")
// 	}
// 	username := cmd.args[0]

// 	err := s.cfg.SetUser(username)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("user set to %s\n", username)

// 	return nil
// }

func (c *commands) run(s *state, cmd command) error {
	if cmdval, ok := c.names[cmd.name]; ok {
		return cmdval(s, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.name)
}

func (c *commands) register(name string, f func(*state, command) error ) {
	c.names[name] = f
}