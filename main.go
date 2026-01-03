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

	myState := state{
		config: cfg,
	}

	myMap := make(map[string]string)
	myCommands := commands{
		mapCommands: myMap,
	}

	myCmd := command{
		name: "login",
		args: os.Args,
	}
	myCommands.register("login", "cmdLogin", myCmd)

}
