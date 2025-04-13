package main

import (
	"fmt"
	"time"
)

func printMessage(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go printMessage("Go is awesome!")
	printMessage("we miss classes")
}
