package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Apple"
		ch <- "Ball"
		close(ch) // Close the channel after sending (optional in this case)
	}()

	// Receive messages after sending
	for i := 0; i < 2; i++ {
		msg := <-ch // Blocks until a value is sent
		fmt.Println(msg)
	}
}
