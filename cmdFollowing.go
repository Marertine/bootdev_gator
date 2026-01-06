package main

import (
	"context"
	"fmt"

	"github.com/Marertine/bootdev_gator/internal/database"
	_ "github.com/lib/pq"
)

func cmdFollowing(s *state, cmd command, user database.User) error {
	myCtx := context.Background()

	//userName := s.cfg.CurrentUserName

	/*myUser, err := s.db.GetUser(myCtx, user.Name)
	if err != nil {
		return err
	}*/

	feedFollows, err := s.db.GetFeedFollowsForUser(myCtx, user.ID)
	if err != nil {
		return err
	}

	for _, ff := range feedFollows {
		fmt.Println(ff.FeedName)
	}

	return nil
}
