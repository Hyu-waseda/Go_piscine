package piscine

import (
	"ft"
	"os"
	"fmt"
)

func Ztail() {
	bytes := Atoi(os.Args[2])
	files := getFiles(os.Args[3:])
	for _, f := range files {
		executeTail(f)
	}
	/* switch strsLen(os.Args) - 1 {
	case 0:
		printlnStr("File name missing")
	case 1:
		catFile(os.Args[1])
	default:
		printlnStr("Too many arguments") 
	}*/
}

func Atoi(s []string) {
	sign := getSign(s)
}

func isDisit(c rune) bool {
	return ('0' <= c && c <= '9')
}

func getFiles(args []string) []string {
	return args[3:]
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
	f, err := os.Open(file)
	if err != nil {
		printlnStr("error")
	}
	defer f.Close()

	var buf []byte
	for {
		n, err := f.Read(buf)
		fmt.Println()
        if n == 0 {
            break
        }
        if err != nil{
            break
        }
		printlnStr(string(buf))
	}
	
}