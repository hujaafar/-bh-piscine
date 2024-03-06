package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	}

	return (arr[n/2-1] + arr[n/2]) / 2
}
