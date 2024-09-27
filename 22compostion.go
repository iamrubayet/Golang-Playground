package main

import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

type contract struct {
	fName string
	lName string
	phone string
}

type business struct {
	name string
	address string
	contact 
}


func (b business) info (){
	pl("",b.name,b.address,b.conact.lname)

}

func main(){

	con1 := customer {
		"abc",
		"tbc",
		"333",
	}

	bus1 := business {
		"sdfsf",
		"asddf",
		con1,


	}

	bus1.info()
	

} 