package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

/*

Client for TCP examples in Go

@author Michael Burke

*/

func DateAtHost(host, port string) error {
	conn, err := net.Dial("tcp", strings.Join([]string(host, port), ":"))
	if err != nil {
		log.Fatal("Error dialing TCP connection to host: ", err)
	}
	defer conn.Close()

	inbuf := make([]byte, 1024)
	time, _ := conn.Read(inbuf)
	fmt.Printf("Received %s", inbuffer[:time])
}

func main() {
	// os.Args[0] is the name of the file
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run DateAtHose.go <hostname> <port>")
		os.Exit(1)
	}
	DateAtHost(os.Args[1], os.Args[2])
}
