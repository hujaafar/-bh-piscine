package piscine

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	items := strings.Fields(str)

	summary := make(map[string]int)

	for _, item := range items {
		summary[item]++
	}

	return summary
}
