package main

import (
	"crypto/sha1"
	"fmt"

	"github.com/asstronom/lab4Signature/permutation"
	"github.com/asstronom/rsa/rsa"
	"go.mongodb.org/mongo-driver/bson"
)

type Server struct {
	con     Connection
	public  rsa.PublicKey
	private rsa.PrivateKey
	boot    chan struct{}
}

func NewServer() (*Server, error) {
	//generate rsa keys for server
	public, private, err := rsa.GenKeys()
	if err != nil {
		return nil, err
	}
	return &Server{
		public:  public,
		private: private,
		boot:    make(chan struct{}, 1),
	}, nil
}

//this function returns Connection type and sets it's own rec and send channels
func (s *Server) Dial() Connection {
	recieve := make(chan []byte)
	send := make(chan []byte)
	s.con = Connection{
		Recieve: recieve,
		Send:    send,
	}
	s.boot <- struct{}{}
	return Connection{
		Recieve: send,
		Send:    recieve,
	}
}

//starts up the server
func (s *Server) Boot() error {
	fmt.Println("server: in sleep mode")
	//wait till recieve "wake up" signal
	<-s.boot
	fmt.Println("server: booted")
	//validate connection
	if s.con.Validate() != nil {
		return s.con.Validate()
	}
	//recieve request from client
	<-s.con.Recieve
	//send public key to client
	bytes, err := bson.Marshal(s.public)
	if err != nil {
		return fmt.Errorf("error marshaling public key %s", err)
	}
	fmt.Printf("server: sent public key\n")
	s.con.Send <- bytes
	//recieve digital signature and document from client
	bytes = <-s.con.Recieve
	ds := DigitalSignature{}
	err = bson.Unmarshal(bytes, &ds)
	if err != nil {
		return fmt.Errorf("error unmarshaling digital signature %s", err)
	}
	fmt.Printf("server: recieved digital signature\n")
	//decrypt symetric key
	bytes = s.private.Decrypt(ds.SymetricKey)
	symetricKey := make([]int, len(bytes))
	for i := range bytes {
		symetricKey[i] = int(bytes[i])
	}
	fmt.Printf("server: decrypted symetric key\nsymetric key: %v\n", symetricKey)
	//decrypt hash
	cipher := permutation.NewPermutationCipher(symetricKey)
	decryptedHash, err := cipher.Decrypt(ds.MessageHash)
	if err != nil {
		return err
	}
	fmt.Printf("server: decrypted hash\ndecrypted hash: %v\n", decryptedHash)
	//hash recieved message
	hasher := sha1.New()
	hasher.Write(ds.Message)
	messageHash := []byte{}
	messageHash = hasher.Sum(messageHash)
	fmt.Printf("server: calculated recieved message hash\nreceived message hash: %v\n", messageHash)
	//check if hashes are the same
	for i, v2 := range messageHash {
		//if hashes are not the same -> return code 1
		if decryptedHash[i] != v2 {
			s.con.Send <- []byte{1}
			return err
		}
	}
	fmt.Printf("server: confirmed that hashes are the same\n")
	//if everything is ok -> return code 0
	s.con.Send <- []byte{}
	return nil
}
