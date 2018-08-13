package user

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

func New(id, age int, first, last string) User {
	return User{ID: id, Age: age, FirstName: first, LastName: last}
}

func (u User) String() string {
	return fmt.Sprintf("%d, %s, %s, %d", u.ID, u.FirstName, u.LastName, u.Age)
}
