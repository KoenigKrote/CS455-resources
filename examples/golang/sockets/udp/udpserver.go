package main

import (
	"log"
	"os"
)

/*

Simple UDP server example in Go

@author Michael Burke

*/

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run udpserver.go <port>")
	}

	port := os.Args[1]

	addr, err := net.ResolveUDPAddR("udp4", ":"+port)
	if err != nil {
		log.Fatal("Error resolving address: ", err)
	}

	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("Error listening on UDP port: ", err)
	}
	defer listener.Close()

	count := 0
	for {
		buf := make([]byte, 1024)

		_, addr, err := listener.ReadFromUDP(buf)
		if err != nil {
			log.Fatal("Error reading UDP Packet: ", err)
		}

		log.Println("Received UDP packet from", addr.String())
		listener.WriteToUDP(buf, addr)
		count += 1
		log.Println("Received", count, "total packets")
	}
}
