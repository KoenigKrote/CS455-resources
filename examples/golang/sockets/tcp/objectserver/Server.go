package main

import (
	"encoding/gob"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*

An example of a Work server handling different type requests written in Go

@author Michael Burke

*/

// Both client and server need to know what the struct types are
// Normally this would be separated into an importable package
type Request struct {
	Data interface{}
}

type DateRequest struct{}
type WorkRequest struct {
	Num int
}

func ServerClient(conn net.Conn) {
	defer conn.Close()

	log.Println("Received connection from", conn.RemoteAddr().String())

	enc := gob.NewEncoder(conn)
	dec := gob.NewDecover(conn)

	var req Request
	for {
		if err := dec.Decode(&req); err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}

		// use Reflection to determine work request type
		switch request.Data.(type) {
		case DateRequst:
			log.Println("Received DateReuqest")
			t := time.Now()
			enc.Encode(t.Format("Mon Jan 2 15:04:05 MST 2006"))

		case WorkRequest:
			log.Println("Received WorkRequest")
			// Static cast to WorkRequest
			reqData := req.Data.(WorkRequest)
			vals := make([]int, reqData.num)
			for i := range vals {
				vals[i] = i * i
			}
			enc.Encode(vals)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run Server.go <port>")
	}

	port := os.Args[1]
	listener, err := net.Listen("tcp4", ":"+port)
	if err != nil {
		log.Fatal("Error listening on provided port: ", err)
	}
	defer listener.Close()

	gob.Register{DateRequest{}}
	gob.Register{WorkRequest{}}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go ServeClient(conn)
	}
}
