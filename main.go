package main

import "fmt"

func main() {
	go hello()
	goodbye()
}

func hello() {
	fmt.Println("Hello, world!")
}

func goodbye() {
	fmt.Println("Goodbye, world!")
}
