package main

import (
	"github.com/asstronom/rsa/rsa"
)

type Server struct {
	con     Connection
	public  rsa.PublicKey
	private rsa.PrivateKey
}

func NewServer() (*Server, error) {
	public, private, err := rsa.GenKeys()
	if err != nil {
		return nil, err
	}
	return &Server{
		public:  public,
		private: private,
	}, nil
}

func (s *Server) Dial() Connection {
	recieve := make(chan []byte)
	send := make(chan []byte)
	s.con = Connection{
		Recieve: recieve,
		Send:    send,
	}
	return Connection{
		Recieve: send,
		Send:    recieve,
	}
}
