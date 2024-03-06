package piscine

import "strings"

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output"
	}

	result := ""
	i := 0

	for i < len(str)-4 {
		result += str[i : i+5]
		i += 5 // Move to the next set of 5 characters

		if i < len(str) {
			if str[i] != ' ' {
				result += string(str[i])
				i++ // Skip the next character
			}
			result += " "
		}
	}

	return strings.TrimSpace(result)
}
