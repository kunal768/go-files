// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	send := make(chan int)
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			send <- i
		}
		close(send)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for {
			val, ok := <-send
			if !ok {
				fmt.Println("breaking")
				break
			}
			fmt.Println("value received from channel", val)
		}
		wg.Done()
	}()

	fmt.Println("waiting for values from channel...")
	wg.Wait()
}
