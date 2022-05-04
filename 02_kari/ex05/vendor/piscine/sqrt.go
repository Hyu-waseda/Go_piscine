package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	for i := 1; ; i++ {
		if i * i == nb {
			return i
		}
		if 2147483647 / i < i {
			return 0
		}
	}
}