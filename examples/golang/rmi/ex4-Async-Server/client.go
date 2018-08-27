package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"../simplermi"
)

func main() {
	client, err := simplermi.SimpleClient(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

	reply := make([]int, 5)
	var args int
	rpcChannel := make(chan *rpc.Call, 5)

	args = 10000
	client.Go("RPCrequest.Calculate", &args, &reply[0], rpcChannel)
	args = 1000
	client.Go("RPCrequest.Calculate", &args, &reply[1], rpcChannel)
	args = 10
	client.Go("RPCrequest.Calculate", &args, &reply[2], rpcChannel)
	args = 50
	client.Go("RPCrequest.Calculate", &args, &reply[3], rpcChannel)
	args = 2
	client.Go("RPCrequest.Calculate", &args, &reply[4], rpcChannel)

	fmt.Println("Synchronizing through channel...")
	for i := 0; i < 5; i++ {
		<-rpcChannel
	}
	fmt.Println("All 5 asynchronous calculation calls have completed")
	fmt.Println(reply)
	client.Close()
}
