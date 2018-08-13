package main

/*

Go does not have 'functional' looking lambdas, but they're effectively the same.
You don't get the cool () -> ((x,y) -> x > y).filter(x -> x != 0).toList(), but you can achieve the same things, arguably more readably

@author Michael Burke

*/

import (
	"fmt"
	"sort"

	// Do NOT use relative pathing in actual Go code
	// These are examples and should not be considered good project structuring
	u "./user"
)

var users = []u.User{
	u.New(1, 10, "Alice", "Vai"),
	u.New(4, 32, "Anna", "Smith"),
	u.New(3, 57, "Steve", "Johnson"),
	u.New(9, 18, "Mike", "Stevens"),
	u.New(10, 24, "Alyssa", "Armstrong"),
	u.New(2, 40, "Jim", "Smith"),
	u.New(8, 34, "Chuck", "Schneider"),
	u.New(5, 22, "Jorje", "Gonzales"),
	u.New(6, 47, "Jane", "Michaels"),
	u.New(7, 60, "Kim", "Berlie"),
}

func main() {
	// Alias sort.Slice to s
	s := sort.Slice
	// Alias a print function for listing results
	p := func(s string, users []u.User) {
		fmt.Println("\n", s)
		for _, u := range users {
			fmt.Println(u)
		}
	}

	s(users, func(i, j int) bool { return users[i].ID > users[j].ID })
	p("By ID:", users)

	s(users, func(i, j int) bool { return users[i].FirstName > users[j].FirstName })
	p("By FirstName:", users)

	s(users, func(i, j int) bool { return users[i].LastName > users[j].LastName })
	p("By LastName:", users)

	s(users, func(i, j int) bool { return users[i].Age > users[j].Age })
	p("By Age:", users)

	// Go does not have built in filters
	var underage []u.User
	for _, u := range users {
		if u.Age < 21 {
			underage = append(underage, u)
		}
	}
	p("Underage Users:", underage)
}
