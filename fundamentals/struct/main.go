package main

import (
	"fmt"
	"howardsolutions/go/fundamentals/server/data"
)

func main() {
	howard := data.Instructor{Id: 4, LastName: "Phung"}
	howard.FirstName = "Howard"
	// print(howard)

	kyle := data.NewInstructor("Kyle", "Simpson")
	print(kyle.Print())

	goCourse := data.Course{Id: 2, Name: "Go Fundamentals", Instructor: howard}

	fmt.Println(goCourse)
}
