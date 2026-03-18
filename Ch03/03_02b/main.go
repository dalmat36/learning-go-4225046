package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p *int = &anInt
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("Value of p:", *p)
	}
	
}
