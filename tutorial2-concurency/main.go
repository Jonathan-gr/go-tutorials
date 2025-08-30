package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	// Goroutine 1
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 21
	}()

	// Goroutine 2
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "message from ch2"
	}()

	// Loop that waits for messages
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
			return // exit after receiving from ch2
		default:
			// runs immediately if no case is ready
			fmt.Println("No messages yet...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}
