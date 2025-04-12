package main

import "fmt"

// global variable
var url = "https://howardphung.com"

func calculateTax(price float32) (float32, float32) {
	return price * 0.05, price * 0.02
}

func birthday(pointerAge *int) {
	*pointerAge++
}

func main() {
	// result1, _ := calculateTax(100)
	// fmt.Println(result1)

	age := 22
	birthday(&age)
	fmt.Println(age)

	// PrintData()
}
