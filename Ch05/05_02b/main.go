package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	dog.Speak()
	dog.Sound = "Arf!"
	dog.Speak()
}

type Dog struct {
   Breed string
   Sound string
}

func (d Dog) Speak() {
	fmt.Printf("The %v says %v!\n", d.Breed, d.Sound)
}
