package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "John", Age: 25}

	fmt.Printf("This is the person %+v \n", person)

	// annonymous struct
	// employee := struct {
	// 	name string
	// 	id   int
	// }{
	// 	name: "Alice",
	// 	id:   345,
	// }

	type Address struct {
		Street string
		City   string
	}

	type Contact struct {
		Name    string
		Address string
		Phone   string
	}

	// contact := Contact{Name: "marc", Address: "Address", Phone: "1223"}

	// fmt.Printf("this is annonymous struct %v \n", contact)

	// POINTER and Struct methods
	fmt.Println("name Before: ", person.Name)
	person.modifyPersonName("howard")
	fmt.Println("name AFter: ", person.Name)

	x := 20
	pointerX := &x
	fmt.Printf("Value of X: %d and address of x %p \n", x, pointerX)

	*pointerX = 30
	fmt.Printf("NEW Value of X: %d and address of x %p \n", x, pointerX)
}

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("Inside scope: ", p.Name)
}
