package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; i < len(arr)-1; i++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[2]
}
