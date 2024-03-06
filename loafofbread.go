package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	i := 0

	for i < len(str)-4 {
		count := 0
		for j := i; j < i+5; j++ {
			if str[j] != ' ' {
				result += string(str[j])
				count++
			}
		}
		if count == 5 {
			i += 6 // Move to the next set of 5 characters
		} else {
			i += count // Skip the current word
		}
		result += "\n" // Add newline after each set of 5 characters
	}

	return result
}
