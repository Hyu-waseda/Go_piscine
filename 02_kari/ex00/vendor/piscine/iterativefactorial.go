package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || 21 <= nb{
		return 0
	} else if nb == 0 {
		return 1
	} else {
		ret := 1
		for i := 1; i <= nb; i++ {
			ret *= i
		}
		return ret
	}
}