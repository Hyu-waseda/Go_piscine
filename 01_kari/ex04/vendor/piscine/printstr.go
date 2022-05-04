package piscine

import "ft"

func PrintStr(s string) {
	for _, v := range []rune(s) {
		ft.PrintRune(v)
	}
}
