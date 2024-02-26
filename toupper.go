package piscine

func ToUpper(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		if s1[i] >= 'a' || s1[i] <= 'z' {
			return string(s1[i] - 32)
		}
	}
}
