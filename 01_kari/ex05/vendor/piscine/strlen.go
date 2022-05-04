package piscine

func StrLen(s string) int {
	i := 0
	for j, _ := range []rune(s) {
		i = j + 1
	}
	return i
}
