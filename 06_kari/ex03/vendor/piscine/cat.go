package piscine

import (
	"ft"
	"os"
)

func Cat() {
	switch strsLen(os.Args) - 1 {
	case 0:
		catStdin()
	default:
		catFiles()
	}
}

func catStdin() {
	arr := [1024]byte{}
	buf := arr[:]
	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			os.Exit(1)
		}
		if n == 0 {
			buf = arr[:]
		}
		printStr(string(buf[:n]))
	}
}

func catFiles() {
	err := false
	for _, file := range os.Args[1:] {
		err = catFile(file) || err
	}
	if err {
		os.Exit(1)
	}
}

func printStr(s string) {
	for _, v := range []rune(s) {
		ft.PrintRune(v)
	}
}
