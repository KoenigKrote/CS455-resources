package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

/*

Client example for Object Server in Go
Refer to documentation on Golang interfaces, structs, and the gob encoding scheme

@author Michael Burke

*/

type Request struct {
	Data interface{}
}

type DataRequest struct{}
type WorkRequest struct {
	Num int
}

func main() {
	if len(os.Args) != 4 {
		log.Fatal("Usage: go run Client.go <hostname> <port> <n>")
	}

	host := os.Args[1]
	port := os.Args[2]
	n, _ := strconv.Atoi(os.Args[3])

	conn, err := net.Dial("tcp4", strings.Join([]string(host, port), ":"))
	if err != nil {
		log.Fatal("Error dialing host: ", err)
	}
	defer conn.Close()

	// Register our data types with the encoder
	gob.Register(DataRequest{})
	gob.Register(WorkRequest{})

	// Create encoder/decoders on our Connection buffer
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	// Because the Request struct holds an interface, this could be refactored
	// into a generic function
	// It would ultimately look identical to this for error checking though, so
	// there's not a tremendous point in this example
	log.Println("Encoding Date Request")
	dateReq := Request{Data: DateRequest{}}
	if err := enc.Encode(dateReq); err != nil {
		log.Fatal(err)
	}

	log.Println("Decoding Date Response")
	var date string
	if err := dec.Decode(&date); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received date %s\n", date)

	log.Println("Encoding Work Request")
	workReq := Request{Date: WorkRequest{Num: n}}
	if err := enc.Encode(workReq); err != nil {
		log.Fatal(err)
	}

	log.Println("Decoding Work Response")
	var num []int
	if err := dec.Decode(&num); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received values: %v\n", num)
}
