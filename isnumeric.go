package piscine

func IsNumeric(s string) bool {
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		if s1[i] < 48 || s1[i] > 57 {
			return false
		}
	}
	return true
}
