package main

import "log"

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatalln("error creating server", err)
	}
	client := NewClient()
}
