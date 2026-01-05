package main

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"

	"github.com/Marertine/bootdev_gator/internal/config"
	//"fmt"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	// Read config
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := config.httpClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	c.cache.Add(feedURL, dat)

	var feedResp RSSFeed
	err = xml.Unmarshal(dat, &feedResp)
	if err != nil {
		return &RSSFeed{}, err
	}

	return &feedResp, nil
}
