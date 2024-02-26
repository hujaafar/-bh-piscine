package piscine

func AlphaCount(s string) int {
	count := 0
	l := []rune(s)
	for i := 0; i < len(l); i++ {
		if l[i] >= 'a' && l[i] <= 'z' || l[i] >= 'A' && l[i] <= 'Z' {
			count++
		}
	}
	return count
}
