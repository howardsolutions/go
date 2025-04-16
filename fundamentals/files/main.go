package main

import (
	"fmt"
	"howardsolutions/go/files/fileutils"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"

	content, err := fileutils.ReadTextFile(filePath)

	if err != nil {
		fmt.Printf("ERROR PANIC!! %v", err)
		newContent := fmt.Sprintf("Orignal: %v\n Double Original: %v%v", content, content, content)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Println(content)
	}

	// Array
	numbers := [5]int{10, 20, 30, 40}
	fmt.Println("This is out array %v\n", len(numbers))

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Printf("This is our matrix: %v\n", matrix)

	// Slices
	// allNumbers := numbers[:]
	fruits := []string{"apple", "banana", "strawberry"}

	fmt.Printf("These are my fruits %v\n", fruits)
	fruits = append(fruits, "orange")
	fmt.Printf("These are my fruits with orange %v\n", fruits)

	moreFruits := []string{"blueberries", "tomato"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("These are my fruits with More fruits %v\n", fruits)
}
