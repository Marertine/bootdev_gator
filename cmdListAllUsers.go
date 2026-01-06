package main

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

func cmdListAllUsers(s *state, cmd command) error {
	myCtx := context.Background()

	respUsers, err := s.db.GetUsers(myCtx)
	if err != nil {
		// All other errors
		return err
	}

	for _, user := range respUsers {
		outputString := user.Name
		currentString := ""
		if outputString == s.cfg.CurrentUserName {
			currentString = " (current)"
		}
		fmt.Printf("* %s%s\n", outputString, currentString)
	}

	return nil
}
