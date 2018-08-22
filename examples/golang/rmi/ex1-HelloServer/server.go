package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	"./shared"
)

func main() {
	port := "5005"
	if len(os.Args) > 2 {
		port = os.Args[1]
	}

	h := new(shared.Hello)

	server := rpc.NewServer()
	if err := server.RegisterName("Hello", h); err != nil {
		log.Fatalln("Could not register RPC server name: ", err)
	}

	listener, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		log.Fatalln("Could not listen on port: ", err)
	}

	server.Accept(listener)
}
