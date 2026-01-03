package main

import (
	"fmt"
	"log"

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

	/*
		// 2. Set user to your name
		if err := cfg.SetUser("michael"); err != nil {
			log.Fatal(err)
		}

		// 3. Read again
		cfg2, err := config.Read()
		if err != nil {
			log.Fatal(err)
		}

		// 4. Print config
		fmt.Println(cfg2)*/
}
