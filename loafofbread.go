package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output"
	}

	var result []rune
	i := 0

	for i < len(str)-4 {
		for j := 0; j < 5; j++ {
			result = append(result, rune(str[i+j]))
		}
		if i < len(str)-5 && str[i+5] != ' ' {
			result = append(result, ' ')
			i++ // Skip the next character
		}
		i += 6 // Move to the next set of 5 characters
	}

	// Trim leading and trailing spaces
	start, end := 0, len(result)
	for start < len(result) && result[start] == ' ' {
		start++
	}
	for end > start && result[end-1] == ' ' {
		end--
	}

	return string(result[start:end])
}
