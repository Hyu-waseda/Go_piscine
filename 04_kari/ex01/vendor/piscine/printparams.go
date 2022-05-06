package piscine

import (
	"ft"
	"os"
)

func PrintParams() {
	for _, s := range os.Args[1:] {
		PrintlnStr(s)
	}
}

func PrintlnStr(s string) {
	for _, c := range []rune(s) {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}