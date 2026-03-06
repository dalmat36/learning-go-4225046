package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("The sum of the integers is:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("The sum of the floats is:", floatSum)

	// No implicit conversion, so we need to convert i1 to float64 before adding it to f2
	total := float64(i1) + f2
	fmt.Println("The total is:", total)

	product := float64(i1) * f2
	fmt.Println("The product is:", product)
}
