package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t)

	n := time.Now()
	fmt.Printf("The time currently is %s\n", n)
	fmt.Printf("This object's type is %T\n", n)
}
