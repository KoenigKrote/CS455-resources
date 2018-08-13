package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

Simple example of storing lambdas and function pointers in Go

@author Michael Burke

*/

func MessageFunc() {
	fmt.Println("\nPointers to regular functions can also be used instead of anonymous ones")
}

func main() {
	// Refer to documentation regarding Go map syntax
	// They are primitive types and are fast, built-in hashmaps

	cmds := make(map[rune]interface{})
	cmds['h'] = func() { fmt.Println("\nType h, s, or q\n") }
	cmds['q'] = func() { os.Exit(0) }
	cmds['s'] = func(s string) { fmt.Println("\nAnonymous functions can accept arguments.", s) }
	cmds['f'] = MessageFunc

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Menu\n-------------\nh) Help\ns) Print 'Some String'\nf) Call function\nq) Quit")
		in, _ := reader.ReadString('\n')
		val := rune(in[0])

		// This works by storing interfaces in our map, and then casting them
		// as functions with appropriate type paramters.
		// interface{} is the closest thing Go has to true generics
		// Use that power VERY selectively, it most cases you should be more specific
		switch val {
		case 'h', 'q', 'f':
			cmds[val].(func())()
		case 's':
			cmds[val].(func(string))("This string was passed as a parameter")
		}
	}
}
