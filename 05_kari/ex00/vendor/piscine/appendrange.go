package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var ret []int
	for i := min; i < max; i++ {
		ret = append(ret, i)
	}
	return ret
}

