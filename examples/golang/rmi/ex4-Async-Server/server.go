package main

import (
	"log"
	"net/rpc"

	"../simplermi"
)

type RPCrequest int

func (r *RPCrequest) Calculate(args *int, reply *int) error {
	// Just useless spin through a for loop to simulate
	// "some request" taking more time than another
	for i := 0; i < *args; i++ {
		*reply = (*args) * (*args)
	}
	return nil
}

func main() {
	rrequest := new(RPCrequest)
	_, listener, err := simplermi.SimpleServer("5005", "RPCrequest", rrequest)
	if err != nil {
		log.Fatalln(err)
	}

	// If we want to handle multiple clients at once, accept clients separately in a for loop
	// This is why we didn't hold on to the returned rpc.Server object in the above SimpleServer() call
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		// Serve the connection in a new goroutine
		go rpc.ServeConn(conn)
	}
}
