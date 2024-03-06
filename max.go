package piscine

func Max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	return max
}
