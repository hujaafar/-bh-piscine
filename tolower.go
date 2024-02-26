package piscine

func ToLower(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		if s1[i] >= 'A' && s1[i] <= 'Z' {
			s1[i] += 32
		}
	}
	return string(s1)
}
