package piscine

func IsNumeric(s string) bool {
	for _, v := range []rune(s) {
		if !('0' <= v && v <= '9') {
			return false
		}
	}
	return true
}
