package main

import "fmt"

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("Calculator 1.0 in GO")
	fmt.Println("=====================")
	fmt.Println(("Which Operations? Add, substract, multiply, divide"))

	fmt.Scanf("%s", &operation)

	fmt.Println("Enter First number")
	fmt.Scanf("%d", &number1)

	fmt.Println("Enter Second number")
	fmt.Scanf("%d", &number2)

	switch operation {
	case "add":
		fmt.Println((number1 + number2))
	case "subtract":
		fmt.Println((number1 - number2))
	case "multiply":
		fmt.Println((number1 * number2))
	case "divide":
		fmt.Println((number1 / number2))
	}
}
