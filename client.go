package main

import "fmt"

type Client struct {
	con Connection
}

func NewClient() *Client {
	client := Client{}
	return &client
}

func (client *Client) ConnectTo(c Connectable) {
	client.con = c.Dial()
}

