package piscine

func StrRev(s string) string {
	ret := ""
	j := 0
	for i := StrLen(s) - 1; i >= 0; i-- {
		ret += string([]rune(s)[i])
		j++
	}
	return ret
}

func StrLen(s string) int {
	i := 0
	for j, _ := range []rune(s) {
		i = j + 1
	}
	return i
}
