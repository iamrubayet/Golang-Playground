package main

import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("Func 1 :", i)
	}
}
func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("Func 2 :", i)
	}
}







func main(){
	go printto15()
	go printto10()
	time.sleep(2*time.second)
	
}

// go routines are unordered by default