package main

import (
	"fmt"
	"errors"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("invalid number of arguments")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("setUser error: %w", err)
	}
	fmt.Println("user has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("invalid number of arguments")
	}
	if s.db[cmd.args[1]]
	s.db.CreateUser(
		ID: uuid.New(),
		CreatedAt: ,
		UpdatedAt,
		Name: cmd.args[1],
	)
}