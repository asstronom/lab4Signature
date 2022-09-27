package main

type Server struct {
	con Connection
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
