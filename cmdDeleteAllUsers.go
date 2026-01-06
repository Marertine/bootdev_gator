package main

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

func cmdDeleteAllUsers(s *state, cmd command) error {
	myCtx := context.Background()

	err := s.db.DeleteAllUsers(myCtx)
	if err != nil {
		// All other errors
		return err
	}

	fmt.Println("Database reset, all users deleted.")

	return nil
}
