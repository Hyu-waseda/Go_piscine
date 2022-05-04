package piscine

import (
	"ft"
	"os"
)

func PrintProgramName() {
	PrintlnStr(os.Args[0][2:])
}

func PrintlnStr(s string) {
	for _, c := range []rune(s) {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}