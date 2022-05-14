package piscine

func IsAlpha(s string) bool {
	for _, v := range []rune(s) {
		if !isAlNum(v) {
			return false
		}
	}
	return true
}

func isAlNum(v rune) bool {
	if ('A' <= v && v <= 'Z') ||
		('a' <= v && v <= 'z') ||
		('0' <= v && v <= '9') {
		return true
	}
	return false
}
