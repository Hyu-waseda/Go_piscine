package piscine

func Split(s , sep string) []string {
	if s == "" && sep == "" {
		return  []string{}
	}
	if StrLen(s) < StrLen(sep) {
		return []string{s}
	}
	ret := []string{}
	for i, _ := range []rune(s) {
		if FindSep(s[i:], sep) {
			if sep == "" {
				ret = append(ret, s[:1])
				return append(ret, Split(s[1:], sep)...)
			} else {
				ret = append(ret, s[:i])
				return append(ret, Split(s[i+StrLen(sep):], sep)...)
			}
		} 
	}
	return []string{s}
}

func FindSep(s, sep string) bool {
	if sep == "" {
		return true
	}
	j := 0
	for i, c := range []rune(s) {
		if i != j {
			return false
		}
		if c == []rune(sep)[j] {
			j++
			if j == StrLen(sep) {
				return true
			}
		}
	}
	return false
}

func StrLen(s string) int {
	len := 0
	for i, _ := range []rune(s) {
		len = i + 1
	}
	return len
}
