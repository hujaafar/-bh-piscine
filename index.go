package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if i+len(toFind) <= len(s) {
			if s[i:(i+len(toFind))] == toFind {
				return i
			}
		}
	}
	return -1
}
