package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kitaclysm/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	// ensure correct number of args and set url variable
	if len(cmd.args) != 1 {
		return errors.New("invalid number of arguments, command requires url")
	}
	url := cmd.args[0]
	
	// get feed
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	// get current user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	// create new follow record
	follows, err := s.db.CreateFeedFollow(context.Background(),database.CreateFeedFollowParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		UserID:		user.ID,
		FeedID:		feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %s\n", follows.FeedName)
	fmt.Printf("User: %s\n", follows.UserName)

	return nil
}