package main

import (
	"errors"
	"fmt"
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

	s.config.CurrentUserName = name
	/*err := s.config.SetUser(name)
	if err != nil {
		return err
	}*/

	fmt.Printf("User has been set as '%s'\n", name)
	return nil
}
