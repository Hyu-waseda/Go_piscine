package piscine

func ConcatParams(args []string) string {
	ret := ""
	for _, s := range(args) {
		ret += s + "\n"
	}
	if ret != "" {
		ret = ret[:StrLen(ret)-1]
	}
	return ret
}

func StrLen(s string) int {
	len := 0
	for i, _ := range s {
		len = i + 1
	}
	return len
}
