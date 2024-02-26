package piscine

func IsUpper(s string) bool {
	s1 := []rune(s)

	found := false
	for index := range s1 {
		if s1[index] >= 'A' && s1[index] <= 'Z' {
			found = true
		} else {
			found = false
		}
	}
	return found
}
