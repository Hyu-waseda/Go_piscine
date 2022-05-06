package main

import (
	//"piscine"
	"os"
	"ft"
)

const (
	yes = 1
	no = 0
	EvenMsg = "I have an even number of arguments"
	OddMsg = "I have an odd number of arguments"
)
var lengthOfArg = int(StrsLen(os.Args) - 1)

type boolean int

func StrsLen(strs []string) int {
	len := 0
	for i, _ := range strs {
		len = i + 1
	}
	return len
}

func even(nbr int) int {
	if nbr % 2 == 0 {
		return 1
	} else {
		return 0
	}
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
