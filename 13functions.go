


package main

import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

func say(){
	pl("dff")
}

func summ(x int, y int) int{
	return x +y

}

func gettwo (x int, y int)(int,int){
	return x+1,x+2
}

func getquotient(x float64, y float64) (ans float64, err error){
	if y == 0 {
		return 0, fmt.printf("you cant divide by zero")

	}else {
		return x/y, nil // success case nil

	}
}


// variadic function means unknow number of values receiving

func getsum2(nums ...int) int {
	sum :=0

	for _,num : range nums {
		sum += num

	}
	return sum
}


func main(){
	// ----- FUNCTIONS -----
	// func funcName(parameters) returnType {BODY}
	// If you only need a function in the current package
	// start with lowercase letter
	// Letters and numbers in camelcase
	sayHello()
	pl(getSum(5, 4))
	f1, f2 := getTwo(5)
	fmt.Printf("%d %d\n", f1, f2)

	// Function that can return an error
	ans, err := getQuotient(5, 0)
	if err == nil {
		pl("5/4 =", ans)
	} else {
		pl(err)
		// End program
		// log.Fatal(err)
	}

	// Function receives unknown number of parameters
	// Variadic Function
	pl("Unknown Sum :", getSum2(1, 2, 3, 4))

	// Pass an array to a function by value
	vArr := []int{1, 2, 3, 4}
	pl("Array Sum :", getArraySum(vArr))

	// Go passes the value to functions so it isn't changed
	// even if the same variable name is used
	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)
	
} 