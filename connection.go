package main

import "fmt"

type Connection struct {
	Recieve <-chan []byte
	Send    chan<- []byte
}

type Connectable interface {
	Dial() Connection
}

func (con *Connection) Validate() error {
	if con.Recieve == nil {
		return fmt.Errorf("error, recieve chan is nil")
	} else if con.Send == nil {
		return fmt.Errorf("error, send chan is nil")
	}
	return nil
}
