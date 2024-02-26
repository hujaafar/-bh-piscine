package piscine

func IsPrintable(s string) bool {
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		if s1[i] < 32 || s1[i] > 126 {
			return false
		}
	}
	return true
}
