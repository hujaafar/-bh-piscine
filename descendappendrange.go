package piscine

func DescendAppendRange(max, min int) []int {
	if min >= max {
		var tab []int = nil
		return tab
	} else {
		tab := []int{}

		for i := min; i < max; i++ {
			tab = append(tab, i)
		}
		return tab
	}
}
