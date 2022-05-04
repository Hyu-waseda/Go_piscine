package piscine

import (
	"ft"
	"os"
)

func SortParams() {
	args := os.Args[1:]
	for i := 0; i < StrsLen(args); i++ {
		for j := i + 1; j < StrsLen(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j] , args[i]
			}
		}
	}
	for _, s := range args {
		PrintlnStr(s)
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
