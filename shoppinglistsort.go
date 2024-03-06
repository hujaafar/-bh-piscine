package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice)-1; j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
