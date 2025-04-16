package main

import "fmt"

// global variable
var url = "https://howardphung.com"

func calculateTax(price float32) (float32, float32) {
	return price * 0.05, price * 0.02
}

func birthday(pointerAge *int) {
	if *pointerAge > 140 {
		panic("Too old to be true")
	}

	*pointerAge++
}

func main() {
	// result1, _ := calculateTax(100)
	// fmt.Println(result1)

	defer fmt.Println("Bye! Done!")

	// age := 22
	// birthday(&age)
	// fmt.Println(age)

	// Enum
	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	fmt.Println("Monday: %d - Tue %d Wed %d", Monday, Tuesday, Wednesday)

	// Enum
	const (
		Jan = iota + 1 // 1
		Feb
		Mar
		Apr
	)

	// PrintData()
}
