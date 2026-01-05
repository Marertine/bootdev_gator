package main

import (
	"context"
	"database/sql"
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

func cmdAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func cmdAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return errors.New("addfeed requires a name and a URL")
	}

	myCtx := context.Background()

	myUser, err := s.db.GetUser(myCtx, s.cfg.CurrentUserName)
	if err != nil {
		// Don't need to test for sql.ErrNoRows separately here because we tested that in login
		// All other errors
		return err
	}

	myFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    myUser.ID,
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

	fmt.Printf("Feed '%s' has been created.\n", feed.Name)
	printFeed(feed)

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID:      %v\n", feed.ID)
	fmt.Printf(" * Name:    %v\n", feed.Name)
	fmt.Printf(" * URL:     %v\n", feed.Url)
	fmt.Printf(" * User ID: %v\n", feed.UserID)
}

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
		fmt.Printf("* %s\n", feed.UserID)
	}

	return nil
}
