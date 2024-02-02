package main

import (
	"fmt"
	"bufio"
	"os"
	"utf8/unicode"
)

var pl = fmt.Println

func main(){
	// character ascii of string

	rStr := "abcd"

	pl("Rune count:",utf8.RuneCountInString(rstr))

	for i,runeVal := range rStr {
		fmt.Printf("%d : %#U : %c",i,runeVal,runeVal)
	} 



	
} 