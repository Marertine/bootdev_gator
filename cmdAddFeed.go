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

func cmdAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("addfeed requires a name and a URL")
	}

	myCtx := context.Background()

	myFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(myCtx, myFeedParams)
	if err != nil {
		/*// Type assertion to *pq.Error
		if pqErr, ok := err.(*pq.Error); ok {
			// Inspect the PostgreSQL error code
			fmt.Println("Postgres error code:", pqErr.Code)
			fmt.Println("Message:", pqErr.Message)
			fmt.Println("Detail:", pqErr.Detail)
			fmt.Println("Constraint:", pqErr.Constraint)

			// Example: unique violation
			if pqErr.Code == "23505" {
				return fmt.Errorf("Feed already exists")
			}
		}*/
		// All other errors
		return err
	}

	_, err = followFeedForUser(myCtx, s.db, user.Name, cmd.Args[1])
	if err != nil {
		return err
	}

	fmt.Printf("Feed '%s' has been created.\n", feed.Name)
	printFeed(feed)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * URL:     %v\n", feed.Url)
	fmt.Printf(" * User ID: %v\n", feed.UserID)
	fmt.Printf(" * LastFetchedAt: %v\n", feed.LastFetchedAt.Time)
}
