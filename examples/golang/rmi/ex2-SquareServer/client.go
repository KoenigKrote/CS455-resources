package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalln("Usage: go run client.go <host> <value(int)> <count> [<port>]")
	}

	host := os.Args[1]
	value, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln("Error converting value to int: ", err)
	}
	count, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalln("Error converting count to int: ", err)
	}

	port := "5005"
	if len(os.Args) > 4 {
		port = os.Args[4]
	}

	client, err := rpc.Dial("tcp4", strings.Join([]string{host, port}, ":"))
	if err != nil {
		log.Fatalln("Could not reach host: ", err)
	}

	var result int
	for i := 0; i <= count; i++ {
		client.Call("Square.Int", value, &result)
		fmt.Printf("Call# %d: result = %d\r", i, result)
	}
	fmt.Println()
}
