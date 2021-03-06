package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	i := nb
	for ; ; i++ {
		if isPrime(i) {
			break ;
		}
	}
	return i
}


func isPrime(nb int) bool {
	if nb == 2 {
		return true
	}
	if nb <= 0 || nb % 2 == 0 {
		return false
	}
	for i := 3; i <= nb / i; i += 2 {
		if nb % i == 0 {
			return false
		}
	}
	return true
}