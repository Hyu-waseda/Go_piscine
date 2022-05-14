package piscine

func ToUpper(s string) string {
	ret := ""
	for _, v := range []rune(s) {
		if 'a' <= v && v <= 'z' {
			ret += string(v + 'A' - 'a')
		} else {
			ret += string(v)
		}
	}
	return ret
}
