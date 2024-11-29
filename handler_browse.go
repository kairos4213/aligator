package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kairos4213/aligator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("%v takes one arg: <number of posts>", cmd.name)
	}
	strLimit := cmd.args[0]
	intLimit, err := strconv.Atoi(strLimit)
	if err != nil {
		return fmt.Errorf("error converting limit to integer: %w", err)
	}

	postsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(intLimit),
	}
	posts, err := s.db.GetPostsForUser(context.Background(), postsParams)
	if err != nil {
		return fmt.Errorf("error getting posts for user: %w", err)
	}

	fmt.Printf("%s's Posts:\n", user.Name)
	fmt.Println("========================")
	for _, post := range posts {
		fmt.Println("----------------------")
		fmt.Printf(" * Feed:        %s\n", post.FeedName)
		fmt.Printf(" * Published:   %s\n", post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Printf(" * Title:       %v\n", post.Title.String)
		fmt.Printf(" * URL:         %s\n", post.Url)
		fmt.Printf(" * Description: %v\n", post.Description.String)
		fmt.Println("----------------------")
	}
	fmt.Println("========================")

	return nil
}
