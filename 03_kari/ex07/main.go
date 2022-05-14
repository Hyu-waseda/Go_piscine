package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
	fmt.Println(piscine.Capitalize("hello! how are you? how+are+things+4you?"))
	fmt.Println(piscine.Capitalize("HELLO! HOW ARE YOU? HOW+ARE+THINGS+4YOU?"))
	fmt.Println(piscine.Capitalize(""))
}
