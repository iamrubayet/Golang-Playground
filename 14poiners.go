
	package main

import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

func changevalue2(myptr *int) int {
	*myptr = 12

}

func dblArrvals(arr *[4]int) int {
	for x:=0;x<4;x++{
		arr[x]*=2
	}

}

func getavg(nums ...float64) float64 {
	
}

func main(){
		// ----- POINTERS -----

	// You can pass by reference with the &
	// (Address of Operator)
	// Print amount and address for amount in memory
	f4 := 10
	pl("f4 :", f4)
	changevalue2(12)
	pl("f4 Address :", &f4)

	// Store a pointer (Pointer to int)
	var f4Ptr *int = &f4
	pl("f4 Address :", f4Ptr)

	// Print value at pointer
	pl("f4 Value :", *f4Ptr)

	// Assign value using pointer
	*f4Ptr = 11
	pl("f4 Value :", *f4Ptr)

	// Change value in function
	pl("f4 before function :", f4)
	changeVal2(&f4)
	pl("f4 after function :", f4)

	// Pass an array by reference
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)

	// Passing a slice to a function works just
	// like when using variadic functions
	// Just add ... after the slice when passing
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))

	
} 