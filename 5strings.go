package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var pl = fmt.Println

func main(){
	sv1 := "a word"
	replacer := strings.NewReplacer("a","another")
	sv2 := replacer.Replace(sv1)
	pl(len(sv2))
	pl("contains another", strings.Contains(sv2,"another"))
	pl("o index", strings.Index(sv2,"o"))
	pl("replace",strings.replace(sv2,"o","0",-1))
	sv3 := "\nsome words\n"
	sv3 = strings.trimspace(sv3)// escaping \t 
	pl("string delimeter",strings.split("a-b-c-d","-"))
	pl("lower:",strings.tolower(sv2))
	pl("Upper",strings.toUpper(sv2))
	pl("prefix",stings.HasPrefix("tocatoc","toca"))// start with that string
	pl("suffix",stings.HasSuffix("tocatoc","toca"))

} 