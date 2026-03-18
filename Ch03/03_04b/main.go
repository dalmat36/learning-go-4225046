package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is an array
	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// This is a slice
	var moreColors = []string{"Yellow", "Purple", "Orange"}
	fmt.Println(moreColors)

	var memoryEfficientColors = make([]string, 0, 3)
	memoryEfficientColors = append(memoryEfficientColors, "Cyan", "Magenta", "Lime")
	fmt.Println(memoryEfficientColors)

	moreColors = remove(moreColors, 1) // Remove "Purple"
	fmt.Println(moreColors)

	sort.Strings(moreColors) // Sort the slice in place
	fmt.Println(moreColors)
}

// remove deletes the element at index i from the slice and returns a new slice
// It combines the elements before and after i, effectively removing the item
func remove(slice []string, i int) []string {
 return append(slice[:i], slice[i+1:]...)
}
