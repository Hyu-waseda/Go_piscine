package piscine

func LastRune(s string) rune {
	var ret rune
	for _, v := range []rune(s) {
		ret = v
	}
	return ret
}