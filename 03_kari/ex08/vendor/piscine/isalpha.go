package piscine

func IsAlpha(s string) bool {
	for _, v := range []rune(s) {
		if IsAlNum(v) == false {
			return false
		}
	}
	return true
}

func IsAlNum(v rune) bool {
	if ('A' <= v && v <= 'Z') ||
		('a' <= v && v <= 'z') ||
		('0' <= v && v <= '9') {
		return true
	}
	return false
}
