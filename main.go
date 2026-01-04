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

	if len(os.Args) < 2 {
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
