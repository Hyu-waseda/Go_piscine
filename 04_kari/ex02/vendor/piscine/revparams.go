package piscine

import (
	"ft"
	"os"
)

func RevParams() {
	for i := StrsLen(os.Args) - 1; i >= 1; i-- {
		PrintlnStr(os.Args[i])
	}
}

func StrsLen(strs []string) int {
	len := 0
	for i, _ := range strs {
		len = i + 1
	}
	return len
}

func PrintlnStr(s string) {
	for _, c := range []rune(s) {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
