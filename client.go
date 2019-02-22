package main

// Client client
type Client struct {
}

// Version app version
const Version string = "1.0.0"

// NewClient create a new client
func NewClient() *Client {
	return &Client{}
}
