package piscine

func IterativeFactorial(nb int) int {
	num := 1

	for i := 1; i <= nb; i++ {
		num = num * nb
	}
	return num
}
