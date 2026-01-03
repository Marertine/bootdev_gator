package config

import (
	"errors"
	"fmt"
	//"math/rand"
	//"os"
	//"time"
)

type command struct {
	name string
	args []string
	//callback    func(*Config, ...string) error
}

type commands struct {
	name        string
	mapCommands map[string]func(*state, command) error
	//callback    func(*Config, ...string) error
}

/* 2. Define a method for the struct
// The (r Rectangle) is the receiver
func (r Rectangle) Area() int {
	return r.width * r.height
}*/

func (c *commands) run(s *state, cmd command) error {
	// Runs the given command with the provided state, if it exists
}

func (c *commands) register(name string, f func(*state, command) error) {
	//This method registers a new handler function for a command name.
}

func cmdLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("'Login' requires a username")
	}

	name := cmd.args[0]

	s.config.CurrentUserName = name
	/*err := s.config.SetUser(name)
	if err != nil {
		return err
	}*/

	fmt.Printf("User has been set as '%s'\n", name)
}
