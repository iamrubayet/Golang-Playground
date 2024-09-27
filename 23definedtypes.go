package main

import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

type Tsp float64
type TBs float64
type ML float64

// Convert with functions (Bad Way)
func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

// Associate method with types
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}





func main(){

	ml1 := ML(TSP(3)*4.92)





	
}