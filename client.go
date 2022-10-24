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
	//validate connection
	if client.con.Validate() != nil {
		return client.con.Validate()
	}
	//send request to server
	client.con.Send <- []byte{}
	//recieve public key from server
	bytes := <-client.con.Recieve
	public := rsa.PublicKey{}
	err := bson.Unmarshal(bytes, &public)
	if err != nil {
		return fmt.Errorf("error unmarshaling public key: %s", err)
	}
	fmt.Println("client: unmarshaled public key")
	//hash message
	hash := make([]byte, 0)
	hasher := sha1.New()
	hasher.Write([]byte(message))
	hash = hasher.Sum(hash)
	fmt.Printf("client: hashed message\nhash: %v\n", hash)
	//generate symetric key
	symetricKey := permutation.GenKey(keyLen)
	fmt.Printf("client: generated symetric key: %v\n", symetricKey)
	symetricCipher := permutation.NewPermutationCipher(symetricKey)
	//encrypt hash
	encryptedHash, err := symetricCipher.Encrypt(hash)
	if err != nil {
		return err
	}
	fmt.Printf("client: encrypted hash with symetric key\nencrypted hash: %v\n", encryptedHash)
	symetricKeyBytes := []byte{}
	for i := range symetricKey {
		symetricKeyBytes = append(symetricKeyBytes, byte(symetricKey[i]))
	}
	//encrypt symetyricKey with rsa public key
	encryptedSymetricKey := public.Encrypt(symetricKeyBytes)
	fmt.Printf("client: encrypted symetric key with public key\n")
	//pack everything
	ds := DigitalSignature{
		SymetricKey: encryptedSymetricKey,
		MessageHash: encryptedHash,
	}
	//if isComprimised, set message to compromised message. server should notice that message is compromised
	if isComprimised {
		ds.Message = []byte(comprimisedMessage)
	} else {
		ds.Message = []byte(message)
	}
	bytes, err = bson.Marshal(ds)
	if err != nil {
		return err
	}
	fmt.Printf("client: sent digital signature\n")
	//get response from server. If it is != 0 -> server noticed that message was compromised
	client.con.Send <- bytes
	if len(<-client.con.Recieve) != 0 {
		return fmt.Errorf("server didn't confirm my message :(")
	}
	return nil
}
