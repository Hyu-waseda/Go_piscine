package piscine

func AtoiBase(s string, base string) int {
	if !validate(base) {
		return 0
	}
	return baseConv(s, base)
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

func baseConv(s string, base string) int {
	ret := 0
	power := firstPower(s, base)
	for _, c := range []rune(s) {
		for i, cBase := range []rune(base) {
			if cBase == c {
				ret += i * power
				power /= strLen(base)
			}
		}
	}
	return ret
}

func firstPower(s string, base string) int {
	ret := 1
	for i := 0; i < strLen(s)-1; i++ {
		ret *= strLen(base)
	}
	return ret
}
