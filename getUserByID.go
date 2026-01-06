package main

import (
	"context"

	"github.com/Marertine/bootdev_gator/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func getUserByID(s *state, userID uuid.UUID) (database.User, error) {
	myCtx := context.Background()
	myUser, err := s.db.GetUserByID(myCtx, userID)
	if err != nil {
		// Don't need to test for sql.ErrNoRows separately here because we tested that in login
		// All other errors
		return database.User{}, err
	}

	return myUser, nil
}
