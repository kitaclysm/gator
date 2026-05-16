package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kitaclysm/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	// ensure correct number of args and set url variable
	if len(cmd.Args) != 1 {
		return errors.New("invalid number of arguments, command requires url")
	}
	url := cmd.Args[0]
	
	// get feed
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	// get current user
	// user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	// if err != nil {
	// 	return err
	// }

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

func handlerListFollows(s *state, cmd command, user database.User) error {
	// ensure correct number of args
	if len(cmd.Args) > 0 {
		return errors.New("command does not accept additional args")
	}

	// get current user
	// user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	// if err != nil {
	// 	return err
	// }

	// get follows for user
	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	// print results
	fmt.Printf("Feeds followed by %s:\n", user.Name)
	for _, follow := range follows {
		fmt.Printf("- %s\n", follow.FeedName)
	}
	return nil
}

func handlerDeleteFollow(s *state, cmd command, user database.User) error {
	// ensure correct number of args
	if len(cmd.Args) != 1 {
		return errors.New("command requires a URL arg")
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.DeleteFollow(context.Background(), database.DeleteFollowParams{
		UserID:	user.ID,
		FeedID:	feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println("Removed feed from user follows")
	return nil
}