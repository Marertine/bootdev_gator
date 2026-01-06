package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

func cmdLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("login requires a username")
	}

	myCtx := context.Background()
	user, err := s.db.GetUser(myCtx, cmd.Args[0])
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 0 rows returned → user not found
			return fmt.Errorf("login: user not found")
		}
		// All other errors
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set as '%s'\n", user.Name)
	return nil
}
