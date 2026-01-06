package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Marertine/bootdev_gator/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func cmdFollowFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("follow requires a url")
	}

	myCtx := context.Background()
	feedURL := cmd.Args[0]
	//userName := s.cfg.CurrentUserName

	followedFeed, err := followFeedForUser(myCtx, s.db, user.Name, feedURL)
	if err != nil {
		return err
	}

	fmt.Printf("Feed '%s' is now followed by user '%s'.\n", followedFeed.FeedName, followedFeed.UserName)
	return nil
}

func followFeedForUser(ctx context.Context, db *database.Queries, userName string, feedURL string) (database.CreateFeedFollowRow, error) {

	feed, err := db.GetFeedByURL(ctx, feedURL)
	if err != nil {
		return database.CreateFeedFollowRow{}, err
	}

	myUser, err := db.GetUser(ctx, userName)
	if err != nil {
		return database.CreateFeedFollowRow{}, err
	}

	myFeedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    myUser.ID,
		FeedID:    feed.ID,
	}

	followedFeed, err := db.CreateFeedFollow(ctx, myFeedFollowParams)
	if err != nil {
		return database.CreateFeedFollowRow{}, err
	}

	return followedFeed, nil

}
