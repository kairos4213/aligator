package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kairos4213/aligator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("%v takes two args: <feed_name> <url>", cmd.name)
	}

	createFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Println("Feed has been created!")
	printFeed(feed)
	fmt.Println("")

	ffParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), ffParams)
	if err != nil {
		return fmt.Errorf("error creating follow record: %w", err)
	}
	fmt.Printf("%s is now following %s\n", user.Name, feedFollow.FeedName)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("%v does not take any args", cmd.name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found")
		return nil
	}

	fmt.Println("Current Feeds: ")
	fmt.Println("-----------------------------------")
	for _, feed := range feeds {
		feedName := feed.Name
		feedUrl := feed.Url
		userName := feed.UserName
		fmt.Printf(" * Feed:  %s\n", feedName)
		fmt.Printf(" * URL:   %s\n", feedUrl)
		fmt.Printf(" * User:  %s\n", userName)
		fmt.Println("")
	}
	fmt.Println("-----------------------------------")

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * id:        %s\n", feed.ID)
	fmt.Printf(" * name:      %s\n", feed.Name)
	fmt.Printf(" * created:   %v\n", feed.CreatedAt)
	fmt.Printf(" * updated:   %v\n", feed.UpdatedAt)
	fmt.Printf(" * url:       %s\n", feed.Url)
	fmt.Printf(" * userID:    %s\n", feed.UserID)
}
