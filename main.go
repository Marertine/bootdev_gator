package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Marertine/bootdev_gator/internal/config"
	"github.com/Marertine/bootdev_gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Boot.Dev/RSS Aggregator Project")

	// Read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Using DB URL:", cfg.DbURL)

	defer db.Close()
	dbQueries := database.New(db)

	myState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	// Initialise the map that will hold the allowed commands
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	// Populate the list of allowed commands
	cmds.register("login", cmdLogin)
	cmds.register("register", cmdRegister)
	cmds.register("reset", cmdDeleteAllUsers)

	// Ensure the appropriate number of command line arguments for each command
	// When compiled...
	// 		os.Args[0] = programname  (eg ./mycli)
	// 		os.Args[1] = command
	// 		os.Args[2] = arg1
	// 		os.Args[3] = arg2
	intRequiredOSArgLength := 0
	switch os.Args[1] {
	case "login":
		intRequiredOSArgLength = 3
	case "register":
		intRequiredOSArgLength = 3
	case "reset":
		intRequiredOSArgLength = 2
	}

	if len(os.Args) < intRequiredOSArgLength {
		log.Fatal("not enough command line arguments")
	}

	myCmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(myState, myCmd)
	if err != nil {
		log.Printf("DEBUG handler error: %v", err)
		log.Fatal(err)
	}

}
