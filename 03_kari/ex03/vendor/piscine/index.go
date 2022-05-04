package piscine

func Index(s string, toFind string) int {
	for i := 0; i + StrLen(toFind) <= StrLen(s); i++ {
		if s[i:i + StrLen(toFind)] == toFind {
			return i
		}
	}
	return -1
}

func StrLen(s string) int {
	i := 0
	for j, _ := range []rune(s) {
		i = j + 1
	}
	return i
}