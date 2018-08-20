package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"./account"
)

func safeRoutine(wg *sync.WaitGroup, acc *account.Account, iterations int) {
	amount := 1.0
	for i := 0; i < iterations; i++ {
		acc.SafeDeposit(amount)
	}
	wg.Done()
}

func unsafeRoutine(wg *sync.WaitGroup, acc *account.Account, iterations int) {
	amount := 1.0
	for i := 0; i < iterations; i++ {
		acc.Deposit(amount)
	}
	wg.Done()
}

func main() {
	if len(os.Args) != 4 {
		log.Fatalln("Usage: go run <goroutines> <iterations> <good|bad>")
		os.Exit(0)
	}

	// Ignore err check on string conversion
	// Don't check for routine ceiling, goroutines are much cheaper than
	// native threads
	numRoutines, _ := strconv.Atoi(os.Args[1])
	iterations, _ := strconv.Atoi(os.Args[2])

	// Declare a function pointer with appropriate signature
	var routine func(*sync.WaitGroup, *account.Account, int)
	// Store which function we want to use based on option
	if strings.EqualFold(os.Args[3], "good") {
		routine = safeRoutine
	} else {
		routine = unsafeRoutine
	}

	// If we want to wait for all goroutines to finish and confirm,
	// then we use WaitGroups.  It's basically a threadpool with built-in
	// conditional variables
	var routinePool sync.WaitGroup
	routinePool.Add(numRoutines)

	a := account.New()
	for i := 0; i < numRoutines; i++ {
		go routine(&routinePool, a, iterations)
	}
	fmt.Println("Waiting for routines to finish...")
	routinePool.Wait()
	fmt.Println("All routines have finished.")
	fmt.Printf("End account value: %.2f\n", a.GetBalance())
}
