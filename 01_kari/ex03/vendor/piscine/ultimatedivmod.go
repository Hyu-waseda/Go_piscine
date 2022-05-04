package piscine

func UltimateDivMod(a *int, b *int) {
	tmp := *a
	*a = *a / *b
	*b = tmp % *b
}
