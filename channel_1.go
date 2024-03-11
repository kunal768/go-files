// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

/*
Here a channel is made in main
Now a go routine created inside main, sends message to the channel 
The 'main' method waits for value from channel hence waiting for the go routine to execute (meaning no need for time.sleep) and then sets the value in msg  
*/

func main() {
	channel := make(chan string)

	go func() {
		channel <- "apple"
	}()

	msg := <-channel

	fmt.Println(msg)

}
