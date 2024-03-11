
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

