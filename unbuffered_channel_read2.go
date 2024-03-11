/*We can also write reading function in go routine and writing function out of go routine 
just to make the process parallel, we can write both in go routine but then use waitGroup to make it synchronous */

package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		// Receive messages after sending
		for i := 0; i < 2; i++ {
			msg := <-ch // Blocks until a value is sent
			fmt.Println(msg)
		}
	}()

	ch <- "Apple"
	ch <- "Ball"
	close(ch) // Close the channel after sending (optional in this case)

}
