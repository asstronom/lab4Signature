package main

import "fmt"


//connection struct that has 2 channels, one for recieve and one for send
type Connection struct {
	Recieve <-chan []byte
	Send    chan<- []byte
}

type Connectable interface {
	Dial() Connection
}

//validates that channels are not nil
func (con *Connection) Validate() error {
	if con.Recieve == nil {
		return fmt.Errorf("error, recieve chan is nil")
	} else if con.Send == nil {
		return fmt.Errorf("error, send chan is nil")
	}
	return nil
}
