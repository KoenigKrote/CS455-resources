package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

/*

Simple UDP socket example in Go

@author Michael Burke

*/

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run udpclient.go <host> <port>")
	}

	host := os.Args[1]
	port := os.Args[2]

	addr, err := net.ResolveUDPAddr("udp4", strings.Join([]string(host, port), ":"))
	if err != nil {
		log.Fatal("Error resolving address: ", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Error dialing host: ", err)
	}
	defer conn.Close()

	if _, err := conn.Write([]byte("Hello World!")); err != nil {
		log.Fatal("Error sending packet: ", err)
	}

	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		log.Fatal("Error reading return packet: ", err)
	}

	fmt.Println(string(buf))
}
