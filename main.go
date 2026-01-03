package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Marertine/bootdev_gator/internal/config"
)

func main() {
	fmt.Println("Boot.Dev/RSS Aggregator Project")

	// 1. Read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Initialise the map that will hold the allowed commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	// Populate the list of allowed commands
	cmds.register("login", cmdLogin)

	myState := &state{
		cfg: &cfg,
	}

	if len(os.Args) < 2 {
		log.Fatal("not enough command line arguments")
	}

	cmdName := os.Args[0]
	cmdArgs := os.Args[1]

	myCmd := command{
		Name: cmdName,
		Args: cmdArgs,
	}

	err = cmds.run(myState, myCmd)
	if err != nil {
		log.Fatal(err)
	}

}
