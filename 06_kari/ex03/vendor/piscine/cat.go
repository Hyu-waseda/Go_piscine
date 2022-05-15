package piscine

import (
	"ft"
	"os"
)

func Cat() {
	switch strsLen(os.Args) - 1 {
	case 0:
		catStdin()
	case 1:
		catFile(os.Args[1])
	default:
		printlnStr("Too many arguments")
	}
}

func catStdin() {
	arr := [2048]byte{}
	buf := arr[:]
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			printlnStr("stdin error")
		}
		if n == 0 {
			buf = arr[:]
		}
		printStr(string(buf[:n]))
	}
}


func printStr(s string) {
	for _, v := range []rune(s) {
		ft.PrintRune(v)
	}
}