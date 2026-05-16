package main

import (
	"fmt"
	"errors"
	"context"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return errors.New("invalid number of arguments")
	}

	err := s.db.ResetDB(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting DB: %w", err)
	}
	fmt.Println("reset database")
	return nil
}