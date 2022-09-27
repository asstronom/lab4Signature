package main

type Client struct {
	Rec  chan []byte
	Send chan []byte
}

func NewClient() *Client {
	client := Client{
		Rec:  make(chan []byte),
		Send: make(chan []byte),
	}
	return &client
}
