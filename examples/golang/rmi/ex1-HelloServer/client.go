package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run client.go <host> [<rpc port>]")
	}

	host := os.Args[1]
	port := "5005" // Go net package uses stringified URIs so don't bother converting
	if len(os.Args) > 2 {
		port = os.Args[2]
	}

	client, err := rpc.Dial("tcp4", strings.Join([]string{host, port}, ":"))
	if err != nil {
		log.Fatalln("Could not reach host: ", err)
	}

	args := "Unused argument"
	var reply string
	err = client.Call("Hello.SayHello", args, &reply)
	if err != nil {
		log.Fatalln("Error making RPC call: ", err)
	}

	fmt.Println("Received Message:", reply)
}
