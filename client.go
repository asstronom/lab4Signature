package main

import (
	"crypto/sha1"
	"fmt"

	"github.com/asstronom/lab4Signature/permutation"
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
	symetricKey := permutation.GenKey(keyLen)
	fmt.Printf("client: generated symetric key: %v\n", symetricKey)
	hash := make([]byte, 0)
	hasher := sha1.New()
	hasher.Write([]byte(message))
	hash = hasher.Sum(hash)
	fmt.Printf("client: hashed message\nhash: %v\n", hash)

	return nil
}
