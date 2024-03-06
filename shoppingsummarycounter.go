package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	var currentWord string

	for _, char := range str {
		if char == ' ' || char == ',' || char == '.' {
			if currentWord != "" {
				summary[currentWord]++
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		summary[currentWord]++
	}

	return summary
}
