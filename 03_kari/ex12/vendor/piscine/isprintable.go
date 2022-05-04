package piscine

func IsPrintable(s string) bool {
	for _, v := range []rune(s) {
		if ' ' <= v && v <= '~' {
			continue
		}
		return false
	}
	return true
}
