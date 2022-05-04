package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return PrintNbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func AtoiBase(s string, base string) int {
	ret := 0
	power := FirstPower(s, base)
	for _, c := range []rune(s) {
		for i, cBase := range []rune(base) {
			if cBase == c {
				ret += i * power
				power /= StrLen(base)
			}
		}
	}
	return ret
}

func StrLen(s string) int {
	len := 0
	for i, _ := range []rune(s) {
		len = i + 1
	}
	return len
}

func FirstPower(s string, base string) int {
	ret := 1
	for i := 0; i < StrLen(s)-1; i++ {
		ret *= StrLen(base)
	}
	return ret
}

func PrintNbrBase(nbr int, base string) string {
	if nbr < StrLen(base) {
		return string([]rune(base)[nbr])
	}
	return PrintNbrBase(nbr/StrLen(base), base) + string([]rune(base)[nbr%StrLen(base)])
}
