package main

// global variable
var url = "https://howardphung.com"

func main() {
	// function scoped
	message := "Hello from go"
	price := 34.4
	print(message, price, url)
}
