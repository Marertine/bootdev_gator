package main

import (
	_ "github.com/lib/pq"
)

func debugcmdTest(s *state, cmd command) error {
	// quick & dirty tester
	scrapeFeeds(s)
	return nil
}
