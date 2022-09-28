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

func (s *Server) Boot() error {
	fmt.Println("server: in sleep mode")
	<-s.boot
	fmt.Println("server: booted")
	if s.con.Validate() != nil {
		return s.con.Validate()
	}
	<-s.con.Recieve
	bytes, err := bson.Marshal(s.public)
	if err != nil {
		return fmt.Errorf("error marshaling public key %s", err)
	}
	fmt.Printf("server: sent public key\n")
	s.con.Send <- bytes
	bytes = <-s.con.Recieve
	ds := DigitalSignature{}
	err = bson.Unmarshal(bytes, &ds)
	if err != nil {
		return fmt.Errorf("error unmarshaling digital signature %s", err)
	}
	fmt.Printf("server: recieved digital signature\n")
	//fmt.Printf("encrpted symetricKeyBytes (len: %d): %v\n", len(ds.SymetricKey), ds.SymetricKey)
	bytes = s.private.Decrypt(ds.SymetricKey)
	//fmt.Printf("symetricKeyBytes (len: %d): %v\n", len(bytes), bytes)
	symetricKey := make([]int, len(bytes))
	for i := range bytes {
		symetricKey[i] = int(bytes[i])
	}
	fmt.Printf("server: decrypted symetric key\nsymetric key: %v\n", symetricKey)
	cipher := permutation.NewPermutationCipher(symetricKey)
	decryptedHash, err := cipher.Decrypt(ds.MessageHash)
	if err != nil {
		return err
	}
	hasher := sha1.New()
	hasher.Write(ds.Message)
	messageHash := []byte{}
	messageHash = hasher.Sum(messageHash)
	for i, v2 := range messageHash {
		if decryptedHash[i] != v2 {
			s.con.Send <- []byte{1}
			return err
		}
	}
	s.con.Send <- []byte{}
	return nil
}
