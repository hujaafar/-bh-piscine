package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	i := 0

	for i < len(str)-4 {
		word := ""
		for j := i; j < i+5; j++ {
			if str[j] != ' ' {
				word += string(str[j])
			} else {
				j++ // Skip the space
			}
		}
		result += word + "\n"
		i += 5 // Move to the next set of 5 characters
	}

	// Add the remaining characters if any
	if i < len(str) {
		result += str[i:] + "\n"
	}

	return result
}
