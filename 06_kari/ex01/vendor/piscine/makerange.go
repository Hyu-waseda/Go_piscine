package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	ret := make([]int, max-min, max-min)
	i := 0
	for n := min; n < max; n, i = n+1, i+1 {
		ret[i] = n
	}
	return ret
}
