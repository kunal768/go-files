// You can edit this code!
// Click here and start typing.


package main

import (
	"fmt"
	"time"
)

func someFunc(num int) {
	fmt.Println(num)
}

func main() {
	go someFunc(1) // child processes forked off of main  (async code)
	someFunc(2)  // synchronous code in main 
	go someFunc(3)  

	time.Sleep(3 * time.Second) // comment this and see the execution, go routines are ignored because we don't wait for them to execute

	fmt.Println("hi")
}
