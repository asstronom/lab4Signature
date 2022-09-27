package main

type Client struct {
	con Connection
}

func NewClient() *Client {
	client := Client{}
	return &client
}
