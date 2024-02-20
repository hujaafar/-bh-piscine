package piscine

func Swap(a *int, b *int) {
	b1 := *b
	a1 := *a
	*a = b1
	*b = a1
}
