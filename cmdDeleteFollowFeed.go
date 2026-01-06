package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Marertine/bootdev_gator/internal/database"
	_ "github.com/lib/pq"
)

func cmdDeleteFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("unfollow requires a url")
	}

	myCtx := context.Background()

	myFeed, err := s.db.GetFeedByURL(myCtx, cmd.Args[0])
	if err != nil {
		return err
	}

	myDeleteParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: myFeed.ID,
	}

	err = s.db.DeleteFeedFollow(myCtx, myDeleteParams)
	if err != nil {
		return err
	}

	fmt.Printf("Feed '%s' is no longer followed by user '%s'.\n", myFeed.Name, user.Name)
	return nil
}
