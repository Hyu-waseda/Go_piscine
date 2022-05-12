package piscine

func RecursiveFactorial(nb int) int {
	intMax := int(^uint(0) >> 1)
	if nb < 0 {
		return 0
	}
	if nb == 0 || nb == 1 {
		return 1
	}
	n := RecursiveFactorial(nb - 1)
	if n == 0 || nb > intMax/n {
		return 0
	}
	return nb * n
}