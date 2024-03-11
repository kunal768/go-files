package main

import "fmt"

func main() {
    ch := make(chan string)
    
    // Receive messages before sending (avoid unbuffered channel deadlock issue)
    for i := 0; i < 2; i++ {
        msg := <-ch  // Blocks until a value is sent
        fmt.Println(msg)
    }
    
    ch <- "Apple"
    ch <- "Ball"
    close(ch)  // Close the channel after sending (optional in this case)
}

/* or rather use a buffered channel */

package main

import "fmt"

func main() {
    ch := make(chan string, 2) // Create a channel with a buffer of size 2

    ch <- "Apple"
    ch <- "Ball"
    close(ch)

    for i := range ch {
        fmt.Println(i)
    }
}

