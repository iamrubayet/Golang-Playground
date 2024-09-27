package main

import (
	"fmt"
	"bufio"
	"os"
)


// These functions will print in order using
// channels
// Func receives a channel and then sends values
// over channels once each time it is called

var pl = fmt.Println

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}
func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func main(){
	channel1 := make(chan int)
	channel2 := make(chan int)

	go nums1(channel1)
	go nums2(channel2)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)



	
} 

// ordering of go routines