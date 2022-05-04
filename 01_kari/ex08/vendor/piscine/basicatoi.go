package piscine

func BasicAtoi(s string) int {
	if Validate(s) == false {
		return 0
	}
	s = s[FirstIndex(s):]
	ret := 0
	power := 1
	for i := StrLen(s) - 1; i >= 0; i-- {
		ret += int([]rune(s)[i]-'0') * power
		power *= 10
	}
	return ret
}

func Validate(s string) bool {
	for _, v := range []rune(s) {
		if '0' <= v && v <= '9' {
			continue
		}
		return false
	}
	return true
}

func FirstIndex(s string) int {
	for i, v := range []rune(s) {
		if v != '0' {
			return i
		}
	}
	return 0
}

func StrLen(s string) int {
	i := 0
	for j, _ := range []rune(s) {
		i = j + 1
	}
	return i
}
