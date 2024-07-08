package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	go func() { messages <- "ping" }()
	messages <- "buffered"

	msg := <-messages
	fmt.Println(msg)
	fmt.Println(<-messages)
}
