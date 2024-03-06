package piscine

func Rot14(s string) string {
	s1 := []rune(s)
	for i, chr := range s1 {
		if chr >= 'a' && chr <= 'z' {
			// Rotate lowercase letters
			s1[i] = ((chr - 'a' + 14) % 26) + 'a'
		} else if chr >= 'A' && chr <= 'Z' {
			// Rotate uppercase letters
			s1[i] = ((chr - 'A' + 14) % 26) + 'A'
		}
	}
	return string(s1)
}
