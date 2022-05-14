package main

import (
	"ft"
	"piscine"
)

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "1234556789")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "0")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "0+12")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(125, "01-2")
	ft.PrintRune('\n')
	piscine.PrintNbrBase(0, "0123456789")
	ft.PrintRune('\n')
}
