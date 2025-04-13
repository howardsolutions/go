package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(200 * time.Millisecond)
	}

	channel <- "Done!"
}

func main() {
	var channel chan string

	go printMessage("Go is awesome!", channel)
	// printMessage("we miss classes")
	response := <-channel

	fmt.Println(response)
}
