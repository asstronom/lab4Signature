package main

type Connection struct {
	Recieve <-chan []byte
	Send    chan<- []byte
}

type Connectable interface {
	SetConnection(con Connection)
}
