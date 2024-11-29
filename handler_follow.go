package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/kairos4213/aligator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("%v takes <url> arg", cmd.name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

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

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("%v takes <url> arg", cmd.name)
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	dfParams := database.DeleteFollowRecordParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = s.db.DeleteFollowRecord(context.Background(), dfParams)
	if err != nil {
		return fmt.Errorf("error deleting follow record: %w", err)
	}

	fmt.Printf("%s is no longer following feed url: %s", user.Name, feed.Name)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("%v does not take any args", cmd.name)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get following list: %w", err)
	}

	if len(following) == 0 {
		fmt.Println("You are not following any feeds")
		return nil
	}

	fmt.Printf("%s is currently following:\n", user.Name)
	fmt.Println("-----------------------------------")
	for _, feed := range following {
		feedName := feed.FeedName
		fmt.Printf(" * %s\n", feedName)
	}
	fmt.Println("-----------------------------------")

	return nil
}
