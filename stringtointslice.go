package piscine

func StringToIntSlice(str string) []int {
	var result []int

	for _, char := range str {
		result = append(result, int(char))
	}

	return result
}
