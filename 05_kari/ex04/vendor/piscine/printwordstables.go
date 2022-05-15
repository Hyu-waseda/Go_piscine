package piscine

import "fmt"

func PrintWordsTables(s []string) {
	fmt.Println(concatParams(s))
}

func concatParams(args []string) string {
	ret := ""
	for _, s := range(args) {
		ret += s + "\n"
	}
	if ret != "" {
		ret = ret[:strLen(ret)-1]
	}
	return ret
}

func SplitWhiteSpaces(s string) []string {
	ret := []string{}
	top := true
	for i, c := range []rune(s) {
		if isWhiteSpace(c) {
			top = true
			continue
		} else if top {
			ret = append(ret, s[i:nextSepIndex(s, i)])
			top = false
		}
	}
	return ret
}

func isWhiteSpace(c rune) bool {
	if c == ' ' || c == '\t' || c == '\n' {
		return true
	}
	return false
}

func nextSepIndex(s string, i int) int {
	for j, c := range []rune(s) {
		if isWhiteSpace(c) && i < j{
			return j
		}
	}
	return strLen(s)
}

func strLen(s string) int {
	len := 0
	for i, _ := range []rune(s) {
		len = i + 1
	}
	return len
}
