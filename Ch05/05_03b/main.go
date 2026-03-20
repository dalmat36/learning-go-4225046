package main

import (
	"fmt"
	"time"
)

//Use go routines when as task would otherwise block the main thread
//working with local files, making web service api requests, etc
func main() {
	go say("Hello from the Goroutine!")
	fmt.Println("Hello from main!")

	//declare anonymous function and run as go routine
	go func(message string) {
		fmt.Println(message)
	}("Hello from the anonymous function!")

	time.Sleep(2 * time.Second)
	fmt.Println("All done!")
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}