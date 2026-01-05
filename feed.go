package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
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
		return nil, err
	}

	// Set a User-Agent header to avoid potential blocking by some servers
	req.Header.Set("User-Agent", "gator")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var feedResp RSSFeed
	err = xml.Unmarshal(dat, &feedResp)
	if err != nil {
		return nil, err
	}

	feedResp.Channel.Title = html.UnescapeString(feedResp.Channel.Title)
	feedResp.Channel.Description = html.UnescapeString(feedResp.Channel.Description)
	for i := range feedResp.Channel.Item {
		feedResp.Channel.Item[i].Title = html.UnescapeString(feedResp.Channel.Item[i].Title)
		feedResp.Channel.Item[i].Description = html.UnescapeString(feedResp.Channel.Item[i].Description)
	}

	return &feedResp, nil
}
