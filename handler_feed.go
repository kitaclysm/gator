package main
import (
	"context"
	"fmt"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kitaclysm/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("invalid number of arguments, command requires name and url")
	}
	
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	
	name	:= cmd.args[0]
	url		:= cmd.args[1]
	
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:			uuid.New(),
		CreatedAt:	time.Now().UTC(),
		UpdatedAt:	time.Now().UTC(),
		Name:		name,
		Url:		url,
		UserID:		user.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", feed)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("command does not accept additional arguments")
	}

	feeds, err := s.db.GetAllFeedsWithUsers(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		name := feed.FeedName
		url := feed.Url
		uname := feed.UserName
		fmt.Printf("NAME: %s, URL: %s, USER: %s\n", name, url, uname)
	}
	return nil
}

// func handlerUsers(s *state, cmd command) error {
// 	if len(cmd.args) > 0 {
// 		return errors.New("invalid number of arguments")
// 	}

// 	users, err := s.db.GetUsers(context.Background())
// 	if err != nil {
// 		return err
// 	}
// 	for _, user := range users {
// 		name := user.Name
// 		if name == s.cfg.CurrentUserName {
// 			fmt.Printf("* %s (current)\n", name)
// 		} else {
// 			fmt.Printf("* %s\n", name)
// 		}
// 	}
// 	return nil
// }