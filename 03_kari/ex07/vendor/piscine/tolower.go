package piscine

func ToLower(s string) string {
	ret := ""
	for _, v := range []rune(s) {
		if 'A' <= v && v <= 'Z' {
			ret += string(v + 'a' - 'A')
		} else {
			ret += string(v)
		}
	}
	return ret
}
