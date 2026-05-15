package main

import (
	"context"
	"fmt"
	"errors"
	"log"
	"time"
)

// func handlerAgg(s *state, cmd command) error {
// 	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Printf("%+v\n", feed)
// 	return nil
// }

func handlerAgg(s *state, cmd command) error {
	// accept arg
	if len(cmd.args) != 1 {
		return errors.New("command requires an additional argument")
	}
	time_between_reqs := cmd.args[0]
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
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	// mark as fetched
	_, err = s.db.MarkFeedFetched(context.Background(), next.ID)
	if err != nil {
		return err
	}
	
	// fetch from url
	feed, err := fetchFeed(context.Background(), next.Url)
	if err != nil {
		return err
	}

	for _, rss := range feed.Channel.Item {
		fmt.Printf("%s\n", rss.Title)
	}
	return nil
}