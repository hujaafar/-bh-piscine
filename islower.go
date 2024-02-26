package piscine

func IsLower(s string) bool {
	s1 := []rune(s)

	for _, ch := range s1 {
		if ch < 'a' || ch > 'z' {
			return false
		}
	}
	return true
}
