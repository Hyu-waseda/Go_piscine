package piscine

func ColLen(board []string) int {
	len := 0
	for i, _ := range board {
		len = i + 1
	}
	return len
}

func RowLen(s string) int {
	len := 0
	for i, _ := range []rune(s) {
		len = i + 1
	}
	return len
}
