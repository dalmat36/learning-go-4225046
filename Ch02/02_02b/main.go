package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42

	fmt.Println(str1, str2, str3)
// 	02_02b: go run main.go
// The quick red fox jumped over the lazy brown dog.

	stringLength, err := fmt.Println("The value is", aNumber)
	if err == nil {
		fmt.Println("String length:", stringLength)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Printf("The value is: %v\n", aNumber)

	fmt.Printf("Data type: %T\n", aNumber)
	// https://pkg.go.dev/fmt@go1.26.1

// 02_02b: go run main.go
// The quick red fox jumped over the lazy brown dog.
// The value is 42
// String length: 16
// The value is: 42
}
