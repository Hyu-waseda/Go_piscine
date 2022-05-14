package piscine

import (
	"ft"
	"os"
)

func PrintProgramName() {
	printlnStr(os.Args[0][2:])
}

func printlnStr(s string) {
	for _, c := range []rune(s) {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}