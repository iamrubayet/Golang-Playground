package main

import (
	"fmt"
	"bufio"
	"os"
)

var pl = fmt.Println

func main(){
	pl("what is your name")
	reader := bufio.NewReader(os.stdin)
	name,err := reader.ReadString('\n')
	if err == nil {
		pl("hello",name)
	}else {
		log.fatal("error occured")
	}

} 