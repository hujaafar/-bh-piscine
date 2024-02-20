package piscine

func UltimateDivMod(a *int, b *int) {
	a1 := *a

	*a = a1 / *b
	*b = a1 % *b
}
