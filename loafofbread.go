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
		if i < len(str)-5 {
			result = append(result, '\n') // Add newline after each set of 5 characters
		}
		i += 6 // Move to the next set of 5 characters
	}

	// Trim trailing newline if present
	if len(result) > 0 && result[len(result)-1] == '\n' {
		result = result[:len(result)-1]
	}

	return string(result)
}
