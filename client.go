package main

type Client struct {
	con Connection
}

func NewClient() *Client {
	client := Client{}
	return &client
}

func (c *Client) SetConnection(con Connection) {
	c.con = con
}
