package piscine

func NRune(s string, n int) rune {
	e := []rune(s)
	if n > len(e) && n < 0 {
		return 0
	} else {
		return e[n]
	}
}
