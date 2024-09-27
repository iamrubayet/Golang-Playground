package main

import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

func main(){
	type Animal interface {
	AngrySound()
	HappySound()
}

// Define type with interface methods and its
// own method
type Cat string

func (c Cat) Attack() {
	pl("Cat Attacks its Prey")
}

// Return the cats name with a type conversion
func (c Cat) Name() string {
	return string(c)
}

//automatically knows not like java or python. automatically knows it want to part of this interface



func (c Cat) AngrySound() {
	pl("Cat says Hissssss")
}
func (c Cat) HappySound() {
	pl("Cat says Purrr")
}
	
} 


var kitty Animal

kitty = Cat("kitty")

kitty.angrysound()

var kitty2 cat = kitty.(cat)

kitty2.attack()