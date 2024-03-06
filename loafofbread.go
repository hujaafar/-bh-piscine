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
		result = append(result, '\n') // Add newline after each set of 5 characters
		i += 6                        // Move to the next set of 5 characters
	}

	// Add the remaining characters if any
	for ; i < len(str); i++ {
		result = append(result, rune(str[i]))
	}

	return string(result)
}
