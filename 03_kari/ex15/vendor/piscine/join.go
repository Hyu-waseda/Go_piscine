package piscine

func Join(strs []string, sep string) string {
	ret := ""
	if StrsLen(strs) == 0 {
		return ret
	}
	for _, s := range strs {
		ret += s + sep
	}
	ret = ret[:StrLen(ret) - StrLen(sep)]
	return ret
}

func StrsLen(strs []string) int {
	len := 0
	for i, _ := range strs {
		len = i + 1
	}
	return len
}

func StrLen(s string) int {
	len := 0
	for i, _ := range []rune(s) {
		len = i + 1
	}
	return len
}
