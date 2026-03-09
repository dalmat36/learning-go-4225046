package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum * 100) / 100
	fmt.Printf("Rounded sum: %v\n\n", sum)

	fmt.Println("The value of PI is:", math.Pi)

	circleRadius := 15.5
	circumference := 2 * math.Pi * circleRadius
	fmt.Printf("Circumference: %.2f\n", circumference)
}
