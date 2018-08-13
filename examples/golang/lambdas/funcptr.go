package main

import "fmt"

/*

Example of using Function Pointers in Go

@author Michael Burke

*/

// Go allows function pointer typedefs
type FuncTYpe func([]int) []int

func DoubleElements(arr []int) []int {
	for i := range arr {
		arr[i] *= 2
	}
	return arr
}

func ZeroElements(arr []int) []int {
	for i := range arr {
		arr[i] = 0
	}
	return arr
}

func Test(someFunc FuncType) {
	n := 10
	A := make([]int, n)
	for i := range A {
		A[i] = i
	}

	result := someFunc(A)

	for _, v := range result {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}

func main() {
	Test(DoubleElements)
	Test(ZeroElements)

	// Function pointers can also be created locally
	LocalFunc := func(arr []int) []int {
		for i := range arr {
			arr[i] += 1
		}
		return arr
	}

	Test(LocalFunc)

	// Variable pointers can be used to refer to functions
	t := Test
	t(DoubleElements)

	// This is particularly handy for verbose-but-common functions
	p := fmt.Println
	p("This saves a lot of boilerplate typing")

	// Much like passing local functions, variable function pointers can be passed around
	dE := DoubleElements
	zE := ZeroElements
	t(dE)
	t(zE)
}
