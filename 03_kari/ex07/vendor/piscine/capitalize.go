package piscine

func Capitalize(s string) string {
	ret := ""
	top := true
	for _, v := range []rune(s) {
		RuneAddBack(&ret, v, &top)
	}
	return ret
}

func RuneAddBack(s *string, v rune, top *bool) {
	if *top {
		if IsAlNum(v) {
			if 'a' <= v && v <= 'z' {
				*s += string(int(v) + 'A' - 'a')
			} else {
				*s += string(v)
			}
			*top = false
		} else {
			*s += string(v)
		}
	} else {
		if IsAlNum(v) {
			if 'A' <= v && v <= 'Z' {
				*s += string(int(v) + 'a' - 'A')
			} else {
				*s += string(v)
			}
		} else {
			*s += string(v)
			*top = true
		}
	}
}

func IsAlNum(v rune) bool {
	if ('A' <= v && v <= 'Z') ||
		('a' <= v && v <= 'z') ||
		('0' <= v && v <= '9') {
		return true
	}
	return false
}
