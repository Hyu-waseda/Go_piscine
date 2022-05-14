package piscine

func NRune(s string, n int) rune {
	for i, v := range []rune(s) {
		if i == n - 1{
			return v
		}
	}
	return 0
}