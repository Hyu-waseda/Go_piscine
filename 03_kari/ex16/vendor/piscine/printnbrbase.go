package piscine

import "ft"

func PrintNbrBase(nbr int, base string) {
	if !validate(base) {
		printStr("NV")
		return
	}
	n := uint(nbr)
	if nbr < 0 {
		n = uint(-1*nbr)
		ft.PrintRune('-')
	}
	printNbrBase2(n, base)
}

func validate(base string) bool {
	if strLen(base) < 2 {
		return false
	} else if haveSign(base) {
		return false
	} else if !isUnique(base) {
		return false
	}
	return true
}

func strLen(s string) int {
	len := 0
	for i, _ := range []rune(s) {
		len = i + 1
	}
	return len
}

func haveSign(s string) bool {
	for _, c := range []rune(s) {
		if c == '+' || c == '-' {
			return true
		}
	}
	return false
}

func isUnique(s string) bool {
	var same int
	for _, c1 := range []rune(s) {
		same = 0
		for _, c2 := range []rune(s) {
			if c1 == c2 {
				same++
			}
		}
		if same != 1 {
			return false
		}
	}
	return true
}

func printStr(s string) {
	for _, v := range []rune(s) {
		ft.PrintRune(v)
	}
}

func printNbrBase2(nbr uint, base string) {
	if nbr < uint(strLen(base)) {
		ft.PrintRune([]rune(base)[nbr])
		return
	}
	printNbrBase2(nbr/uint(strLen(base)), base)
	nbr %= uint(strLen(base))
	ft.PrintRune([]rune(base)[nbr])
}
