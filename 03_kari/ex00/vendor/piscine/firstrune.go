package piscine

func FirstRune(s string) rune {
	for i, _ := range []rune(s) {
		if i == 0 {
			return []rune(s)[0]
		} 
	}
	return 0
}