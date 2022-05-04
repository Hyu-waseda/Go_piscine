package piscine

func Swap(a *int, b *int) {
	tmp := *b
	*b = *a
	*a = tmp
}
