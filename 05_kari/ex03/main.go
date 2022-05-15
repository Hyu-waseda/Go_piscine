package main

import (
	"piscine"
	"fmt"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello \t how\n are \n\tyou?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" \t  \n\n\t  \t\n  "))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(""))
}
