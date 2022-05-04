package piscine

func LastRune(s string) rune {
	var r rune
	for _, v := range []rune(s) {
		r = v
	}
	return r
}