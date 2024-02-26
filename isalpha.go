package piscine

func IsAlpha(s string) bool {
	s1 := []rune(s)

	for _, ch := range s1 {
		if ch < 'a' && ch > 'z' || ch < 'A' && ch > 'Z' || ch < '0' && ch > '9' {
			return false
		}
	}
	return true
}
