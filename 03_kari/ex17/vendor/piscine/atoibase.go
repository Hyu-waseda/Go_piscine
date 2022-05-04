package piscine

func AtoiBase(s string, base string) int {
	if !Validate(base) {
		return 0
	}
	return BaseConv(s, base)
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

func BaseConv(s string, base string) int {
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

func FirstPower(s string, base string) int {
	ret := 1
	for i := 0; i < StrLen(s)-1; i++ {
		ret *= StrLen(base)
	}
	return ret
}
