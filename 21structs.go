
package main

import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

type customer struct {
	name string
	address string
	balance float64
}

// normal function manipulating the structs


func getCustomerInfo(c customer){
	pl("%s owes us %.2f\n",c.name,c.balance)

}

func addCustomerAddress(c *customer,address string){
	c.address = address


}


// complicated struct

type rectangle struct {
	length,height float64
}


// now associated fucntion with struct it is not only work but also part of the struct


func (r rectangle) Area() float64 {
	return r.length * r.height

}






func main(){
	 var ts customer
	 ts.name =  "Tom smith"
	 ts.address = "goran"
	 ts.balance =  234.5
	 getCustomerInfo(ts)
	 addCustomerAddress(&ts,"12 th street")
	 pl("address",ts.address)

	 // struct literal another way

	 ss:= customer {"sally","12 main",0.0}
	 pl("ss name", ss.name)

	 rect1 := rectangle {10,20}

	 pl("rect area":react1.Area())








	
}