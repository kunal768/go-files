package main

import (
	"fmt"
)

/* Here int gets printed as it is processed faster by golang */
func main() {
	stringChannel := make(chan string)
	intChannel := make(chan int)

	go func() {
		stringChannel <- "apple"
	}()

	go func() {
		intChannel <- 5
	}()

	select {
	case stringValue := <-stringChannel:
		fmt.Println(stringValue)
	case intValue := <-intChannel:
		fmt.Println(intValue)
	}
}
