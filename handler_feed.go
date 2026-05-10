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