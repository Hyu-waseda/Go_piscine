package piscine

import "ft"

func PrintNbrBase(nbr int, base string) {
	if !Validate(base) {
		PrintStr("NV")
		return
	}
	if nbr < 0 {
		nbr *= -1
		ft.PrintRune('-')
	}
	PrintNbrBase2(nbr, base)
}

func Validate(base string) bool {
	if StrLen(base) < 2 {
		return false
	} else if HaveSign(base) {
		return false
	} else if !IsUnique(base) {
		return false
	}
	return true
}

func StrLen(s string) int {
	len := 0
	for i, _ := range []rune(s) {
		len = i + 1
	}
	return len
}

func HaveSign(s string) bool {
	for _, c := range []rune(s) {
		if c == '+' || c == '-' {
			return true
		}
	}
	return false
}

func IsUnique(s string) bool {
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

func PrintNbrBase2(nbr int, base string) {
	if nbr < StrLen(base) {
		ft.PrintRune([]rune(base)[nbr])
		return
	}
	PrintNbrBase2(nbr/StrLen(base), base)
	nbr %= StrLen(base)
	ft.PrintRune([]rune(base)[nbr])
}

func PrintStr(s string) {
	for _, v := range []rune(s) {
		ft.PrintRune(v)
	}
}
