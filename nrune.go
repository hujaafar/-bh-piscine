package piscine

func NRune(s string, n int) rune {
	e := []rune(s)
	if n > len(e) {
		return 0
	} else if n < 0 {
		n *= -n
		return rune(s[len(s)-n])
	} else {
		return e[n]
	}
}
