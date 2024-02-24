package piscine

func IterativeFactorial(nb int) int {
	num := 1

	for i := 1; i <= nb; i++ {
		if nb == 0 {
			return 0
		}
		num = num * i
	}
	return num
}
