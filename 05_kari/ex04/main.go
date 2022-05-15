package main

import "piscine"

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	piscine.PrintWordsTables(a)
	a = piscine.SplitWhiteSpaces("")
	piscine.PrintWordsTables(a)
}
