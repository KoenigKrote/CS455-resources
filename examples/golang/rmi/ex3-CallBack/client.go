package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"time"

	"../simplermi"
)

func main() {
	client, err := simplermi.SimpleClient(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	// Synchronous call
	var strReply string
	args := "args"
	if err := client.Call("RPCrequest.GetDate", &args, &strReply); err != nil {
		log.Fatalln("Error requesting date: ", err)
	}
	fmt.Println(strReply)

	// Asynchronous call
	// client.Go takes a 4th parameter as a channel of type *rpc.Call
	// This channel will return the call structure representing the invoked function
	// The channel behaves as a "finished" signal
	asyncChannel := make(chan *rpc.Call, 1)
	client.Go("RPCrequest.GetDate", &args, &strReply, asyncChannel)

	fmt.Println("Sleeping after request is made...")
	time.Sleep(250)
	fmt.Println("Waking up")

	// Read in from the channel
	// This will be a blocking call if the channel is empty
	// We also don't care about what we actually get, so don't store it in a variable
	<-asyncChannel
	fmt.Printf("%v\n", strReply)

	//client.Call()
}
