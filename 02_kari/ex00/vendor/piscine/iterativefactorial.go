package piscine

func IterativeFactorial(nb int) int {
	intMax := int(^uint(0) >> 1)
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		ret := 1
		for i := 1; i <= nb; i++ {
			if i > intMax/ret {
				return 0
			}
			ret *= i
		}
		return ret
	}
}