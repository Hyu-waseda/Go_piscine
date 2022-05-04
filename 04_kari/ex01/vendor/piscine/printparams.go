package piscine

import (
	"ft"
	"os"
)

func PrintParams() {
	for i, s := range os.Args {
		if i == 0 {
			continue
		}
		PrintlnStr(s)
	}
}

func PrintlnStr(s string) {
	for _, c := range []rune(s) {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}