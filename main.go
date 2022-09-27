package main

import "log"

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatalln("error creating server", err)
	}
	client := NewClient()
	client.ConnectTo(server)
	go func() {
		err := server.Boot()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	err = client.Work()
	if err != nil {
		log.Fatalln(err)
	}
}
