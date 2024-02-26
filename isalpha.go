package piscine

func IsAlpha(s string) bool {
	s1 := []rune(s)

	for index := range s1 {
		if index >= 20 && index <= 47 {
			return false
		}
	}
	return true
}
