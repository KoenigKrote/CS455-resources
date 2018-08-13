package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"time"
)

/*

Multi Threaded time server example written in Go

@author Michael Burke

*/

func ServeClient(conn net.Conn) {
	log.Println("Received connection from", conn.RemoteAddr().String())

	t := time.Now()
	out := bufio.NewWriter(conn)

	// Refer to Golang documentation on Time Formatting
	out.Write([]byte(t.Format("Mon Jan 2 15:04:05 MST 2009")))
	out.Flush()
	conn.Close()
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run TimeServer.go <port>")
	}

	port := os.Args[1]
	listener, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		log.Fatal("Error listening on port ", port, ", may be in use: ", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		// go keyword spawns a new goroutine
		go ServeClient(conn)
	}
}
