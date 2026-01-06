package main

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

func debugcmdTest(s *state, cmd command) error {
	// quick & dirty test command to test fetchFeed
	ctx := context.Background()

	feed, err := fetchFeed(ctx, "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", feed)
	return nil
}
