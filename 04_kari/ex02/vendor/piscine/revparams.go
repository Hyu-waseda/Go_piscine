package piscine

import (
	"ft"
	"os"
)

func RevParams() {
	for i := strsLen(os.Args) - 1; i >= 1; i-- {
		printlnStr(os.Args[i])
	}
}

func strsLen(strs []string) int {
	len := 0
	for i, _ := range strs {
		len = i + 1
	}
	return len
}

func printlnStr(s string) {
	for _, c := range []rune(s) {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
