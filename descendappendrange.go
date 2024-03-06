package piscine

func DescendAppendRange(max, min int) []int {
	if min >= max {
		return []int{}
	} else {
		tab := []int{}

		for i := max; i > min; i-- {
			tab = append(tab, i)
		}
		return tab
	}
}
