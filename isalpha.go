package piscine

func IsAlpha(s string) bool {
	s1 := []rune(s)

	for i := 0; i < len(s1); i++ {
		if s1[i] >= 20 && s1[i] <= 47 {
			return false
		}
	}
	return true
}
