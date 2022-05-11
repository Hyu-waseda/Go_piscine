package piscine

func IsPrime(nb int) bool {
	if nb == 2 {
		return true
	}
	if nb <= 1 || nb % 2 == 0 {
		return false
	}
	for i := 3; i <= nb / i; i += 2 {
		if nb % i == 0 {
			return false
		}
	}
	return true
}