package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	item := ""
	for _, char := range str {
		if char == ' ' {
			summary[item]++
			item = ""
		} else {
			item += string(char)
		}
	}
	summary[item]++
	return summary
}
