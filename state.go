package main

import (
	"github.com/Marertine/bootdev_gator/internal/config"
	"github.com/Marertine/bootdev_gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
