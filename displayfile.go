package main

import "os"

func main() {
	displayFile()
}

func displayFile() {
	switch StrsLen(os.Args) - 1 {
	case 0:
		PrintlnStr("File name missing")

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
	for _, v := range []rune(s) {
		ft.PrintRune(v)
	}
	ft.PrintRune("¥n")
}