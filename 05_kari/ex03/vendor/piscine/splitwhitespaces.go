package piscine

func SplitWhiteSpaces(s string) []string {
	ret := []string{}
	top := true
	for i, c := range []rune(s) {
		if IsWhiteSpace(c) {
			top = true
			continue
		} else if top {
			ret = append(ret, s[i:NextSepIndex(s, i)])
			top = false
		}
	}
	return ret
}

func IsWhiteSpace(c rune) bool {
	if c == ' ' || c == '\t' || c == '\n' {
		return true
	}
	return false
}

func NextSepIndex(s string, i int) int {
	for j, c := range []rune(s) {
		if IsWhiteSpace(c) && i < j{
			return j
		}
	}
	return StrLen(s)
}

func StrLen(s string) int {
	len := 0
	for i, _ := range []rune(s) {
		len = i + 1
	}
	return len
}
