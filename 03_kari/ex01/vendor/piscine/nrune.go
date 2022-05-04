package piscine

func NRune(s string, n int) rune {
	for i, v := range []rune(s) {
		if i + 1 == n {
			return v
		}
	}
	return 0
}