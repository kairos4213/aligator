package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
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

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Printf("error getting next feed to fetch: %v\n", err)
		return
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed as fetched: %v\n", err)
		return
	}

	feedData, err := fetchRSSFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("error fetching feed data: %v\n", err)
		return
	}

	for _, item := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		postParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       sql.NullString{String: item.Title, Valid: true},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		}
		_, err := s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			if strings.Contains(
				err.Error(),
				"duplicate key value violates unique constraint",
			) {
				continue
			}
			log.Printf("Couldn't create post: %v\n", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
