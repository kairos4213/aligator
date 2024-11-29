package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/kairos4213/aligator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("%v takes one arg: <durationString>", cmd.name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing time duration: %w", err)
	}
	fmt.Printf("Collecting feeds every %v\n\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to fetch: %w", err)
	}

	nextFeedParams := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now().UTC(),
		LastFetchedAt: sql.NullTime{Time: time.Now().UTC()},
		ID:            nextFeed.ID,
	}
	err = s.db.MarkFeedFetched(context.Background(), nextFeedParams)
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	feed, err := fetchRSSFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Println(feed.Channel.Title)
	fmt.Println("=========================================================")
	for _, item := range feed.Channel.Item {
		fmt.Printf("Item Title: %s\n", item.Title)
	}
	fmt.Println("=========================================================")
	return nil
}
