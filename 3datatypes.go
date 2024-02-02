package main

import (
	"fmt"
	"bufio"
	"os"
	"reflect"
	"strcnv"
)


var pl = fmt.Println

func main(){
	//int , float64 , bool
	// these things will have default value
	pl(refelect.Typeof("hello")) // knowing the type
	cv1 := 1.2
	cv2 := int(cv1)// type casting
	cv3 := "50000"
	cv4,err := strconv.Atoi(cv3) //ascii to interger // error handling

	cv5 := 5000
	cv6,_ := strconv.Itoa(cv5) // interger to ascii // blank to identifier

	cv7 := "3.14" // string to float

	if cv8,err := strconv.parsefloat(cv7,64); err==nil {
		pl(cv8)
	}

	cv9 := fmt.Sprintf("%f",3.14) // float to string



	
}