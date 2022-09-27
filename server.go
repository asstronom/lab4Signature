package main

import (
	"fmt"

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
		boot: make(chan struct{})
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
	<-s.boot
	fmt.Println("server booted")
	if s.con.Validate() != nil {
		return s.con.Validate()
	}
	<-s.con.Recieve
	bytes, err := bson.Marshal(s.public)
	if err != nil {
		return fmt.Errorf("error marshaling public key %s", err)
	}
	s.con.Send <- bytes
	bytes = <-s.con.Recieve
	if bytes == nil {
		return fmt.Errorf("error, recieved nil")
	}
	return nil
}
