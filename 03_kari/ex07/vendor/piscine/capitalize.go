package piscine

func Capitalize(s string) string {
	ret := ""
	top := true
	for _, v := range []rune(s) {
		runeAddBack(&ret, v, &top)
	}
	return ret
}

func runeAddBack(s *string, v rune, top *bool) {
	if *top {
		*s += ToUpper(string(v))
	} else {
		*s += ToLower(string(v))
	}
	if *top == true && isAlNum(v) {
		*top = false
	} else if *top == false && !isAlNum(v) {
		*top = true
	}
}

func isAlNum(v rune) bool {
	if ('A' <= v && v <= 'Z') ||
		('a' <= v && v <= 'z') ||
		('0' <= v && v <= '9') {
		return true
	}
	return false
}
