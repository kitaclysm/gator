package main

import (
	"context"
	"fmt"
	"errors"
	"strconv"

	"github.com/kitaclysm/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		converted, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return errors.New("Invalid argument, command requires an integer")
		}
		limit = converted
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID:	user.ID,
		Limit:	int32(limit),
	})
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("%+v\n", post)
	}
	return nil
}