package main

import (
	"fmt"

	"github.com/asstronom/rsa/rsa"
	"go.mongodb.org/mongo-driver/bson"
)

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

func (client *Client) Work() error {
	if client.con.Validate() != nil {
		return client.con.Validate()
	}
	client.con.Send <- []byte{}
	bytes := <-client.con.Recieve
	public := rsa.PublicKey{}
	err := bson.Unmarshal(bytes, &public)
	if err != nil {
		return fmt.Errorf("error unmarshaling public key: %s", err)
	}
	fmt.Println("client: unmarshaled public key")
	return nil
}
