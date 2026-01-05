package main

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	c := Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}

	return c
}
