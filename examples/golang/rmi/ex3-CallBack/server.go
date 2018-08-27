package main

import (
	"log"
	"os"
	"time"

	"../simplermi"
)

// The types on these isn't particularly relevant
type RPCrequest int

func (r *RPCrequest) GetDate(args *string, reply *string) error {
	t := time.Now()
	*reply = t.Format("Mon Jan 2 2006")
	return nil
}

func main() {
	port := "5005"
	if len(os.Args) > 2 {
		port = os.Args[1]
	}

	rrequest := new(RPCrequest)

	server, listener, err := simplermi.SimpleServer(port, "RPCrequest", rrequest)
	if err != nil {
		log.Fatalln("Error setting up server: ", err)
	}

	server.Accept(listener)
}
