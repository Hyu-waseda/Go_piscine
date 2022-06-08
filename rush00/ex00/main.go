package main

import (
	"ft"
	"piscine"
)

func main() {
	board := []string{}
	

	Checkmate(board)
}

func Checkmate(board []string) {
	if !piscine.Validate(board) {
		printlnStr("Input error")
		return
	}
	if piscine.Checkmate(board) {
		printlnStr("Success")
	} else {
		printlnStr("Fail")
	}
}

func printlnStr(str string) {
	for _, c := range str {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
