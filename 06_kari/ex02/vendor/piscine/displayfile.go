package piscine

import (
	"ft"
	"os"
)

func DisplayFile() {
	switch strsLen(os.Args) - 1 {
	case 0:
		printlnStr("File name missing")
	case 1:
		catFile(os.Args[1])
	default:
		printlnStr("Too many arguments")
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
	for _, v := range []rune(s) {
		ft.PrintRune(v)
	}
	ft.PrintRune('\n')
}

func catFile(file string) {
	f, err := os.ReadFile(file)
	if err != nil {
		printlnStr("ERROR: open " + file + ": no such file or directory")
		return
	}
	os.Stdout.Write(f)
	ft.PrintRune('\n')
}