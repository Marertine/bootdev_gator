package main

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

func cmdListAllFeeds(s *state, cmd command) error {
	myCtx := context.Background()

	respFeeds, err := s.db.GetFeeds(myCtx)
	if err != nil {
		// All other errors
		return err
	}

	for _, feed := range respFeeds {
		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("* %s\n", feed.Url)
		myUser, err := getUserByID(s, feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("* %s\n", myUser.Name)
	}

	return nil
}
