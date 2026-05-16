package main

import (
	"context"
	"fmt"
	"errors"
	"log"
	"time"
	"strings"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kitaclysm/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	// accept arg
	if len(cmd.Args) != 1 {
		return errors.New("command requires an additional argument")
	}
	time_between_reqs := cmd.Args[0]
	t, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %s", t)

	ticker := time.NewTicker(t)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			log.Println("error scraping feeds:", err)
		}
	}
	return nil
}

func scrapeFeeds(s *state) error {
	// get next feed
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	// mark as fetched
	_, err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}
	
	// fetch from url
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	for _, item := range rssFeed.Channel.Item {
		// check/parse description
		descr := sql.NullString{}
		if item.Description != "" {
			descr = sql.NullString{
				String:	item.Description,
				Valid:	true,
			}
		}

		// check/parse publish time
		publishedAt := sql.NullTime{}
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			t, err = time.Parse(time.RFC1123, item.PubDate)
		}
		if err == nil {
			publishedAt = sql.NullTime{
				Time:	t,
				Valid:	true,
			}
		} 

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:				uuid.New(),
			CreatedAt:		time.Now().UTC(),
			UpdatedAt:		time.Now().UTC(),
			Title:			item.Title,
			Url:			item.Link,
			Description:	descr,
			PublishedAt:	publishedAt,
			FeedID:			nextFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Println("error creating post: ", err)
		}
	}
	return nil
}