package main

import (
	"log"
	"net"
)

func main() {

	//region Main Server Initialization
	s := newServer()
	go s.run()
	//endregion

	//region TCP Server Initialization
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Started server on port :8888")
	//endregion

	//region Accepting new clients
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
	//endregion
}
