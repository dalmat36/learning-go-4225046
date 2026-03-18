package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["NY"] = "New York"
	states["CA"] = "California"
	states["TX"] = "Texas"

	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)
	delete(states, "TX")
	fmt.Println(states)

	//print keys and values unordered
	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// To get the keys in a specific order, we can create a slice of keys and sort it
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted keys:")
	// Now we can iterate over the sorted keys to access the values in a specific order
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
