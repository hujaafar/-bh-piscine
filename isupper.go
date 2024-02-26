package piscine

func IsUpper(s string) bool {
	s1 := []rune(s)

	for _, ch := range s1 {
		if ch < 'A' || ch > 'Z' {
			return false
		}
	}
	return true
}
