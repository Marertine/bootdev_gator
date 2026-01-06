package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Marertine/bootdev_gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func cmdRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("register requires a username")
	}

	myCtx := context.Background()
	myUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	}

	user, err := s.db.CreateUser(myCtx, myUserParams)
	if err != nil {
		// Type assertion to *pq.Error
		if pqErr, ok := err.(*pq.Error); ok {
			// Inspect the PostgreSQL error code
			fmt.Println("Postgres error code:", pqErr.Code)
			fmt.Println("Message:", pqErr.Message)
			fmt.Println("Detail:", pqErr.Detail)
			fmt.Println("Constraint:", pqErr.Constraint)

			// Example: unique violation
			if pqErr.Code == "23505" {
				return fmt.Errorf("User already exists")
			}
		}
		// All other errors
		return err
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User '%s' has been created.\n", user.Name)
	printUser(user)

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
