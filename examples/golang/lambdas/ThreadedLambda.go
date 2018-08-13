package main

import (
	"fmt"
	"time"
)

/*

Go doesn't have to unwind everything through anonymous classes and instantiation like Java does
As such, threading a lambda is identical to threading any other kind of function

@author Michael Burke

*/

func main() {
	lambdaFunc := func() { fmt.Println("Lambda Function!") }
	go lambdaFunc()
	go func() { fmt.Println("Threaded Lambda without a pointer!") }()
	go func(m string) { fmt.Println(m) }("Threaded lambda with an argument!")
	time.Sleep(250)
}
