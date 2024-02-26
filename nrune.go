package piscine

func NRune(s string, n int) rune {
	e := []rune(s)
	if n < 0 || n > len(e) {
		return 0
	} else {
		return e[n-1]
	}
}
