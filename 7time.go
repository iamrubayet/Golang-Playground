package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
)

var pl = fmt.Println

func main(){
		now := time.Now()
	pl(now.year(),now.day(),now.month())
	pl(now.hour(),now.minute(),now.second())
	
} 