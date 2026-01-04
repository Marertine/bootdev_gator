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

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	//This method registers a new handler function for a command name.
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	// Runs the given command with the provided state, if it exists
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

func cmdLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("'Login' requires a username")
	}

	name := cmd.Args[0]

	s.cfg.CurrentUserName = name
	err := s.cfg.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set as '%s'\n", name)
	return nil
}

func cmdRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("'Register' requires a username")
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

	fmt.Printf("User '%s' has been created.", user.Name)
	fmt.Printf("DEBUG: %v user: %s inserted with UUID: %v\n", user.CreatedAt, user.Name, user.ID)

	return nil
}
