package main

import (
	"log"
	"net"
	"net/rpc"
	"os"
)

// RPC calls and types don't necessarily have to be in their own package
// Though this is usually beneficial to provide a shared Arg/Reply struct
// pairs for RPC calls that want more than one argument and return
type Square int

func (s *Square) Int(args *int, reply *int) error {
	*reply = (*args) * (*args)
	return nil
}

// Golang does not have generics or overloads
// If we wanted to run this with floats, we'd need a separate call
func (s *Square) Float(args *float32, reply *float32) error {
	*reply = (*args) * (*args)
	return nil
}

func main() {
	port := "5005"
	if len(os.Args) > 2 {
		port = os.Args[1]
	}

	s := new(Square)
	server := rpc.NewServer()
	if err := server.RegisterName("Square", s); err != nil {
		log.Fatalln("Could not register RPC server name: ", err)
	}

	listener, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		log.Fatalln("Could not listen on port: ", err)
	}

	server.Accept(listener)
}
