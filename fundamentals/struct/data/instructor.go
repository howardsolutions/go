package data

import "fmt"

type Instructor struct {
	Id        int
	FirstName string
	LastName  string
	Score     int
}

func NewInstructor(firstName string, lastName string) Instructor {
	return Instructor{FirstName: firstName, LastName: lastName}
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d) \n", i.LastName, i.FirstName, i.Score)
}
