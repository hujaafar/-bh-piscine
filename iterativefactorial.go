package piscine

func IterativeFactorial(nb int) int {
	num := 1
	if nb <= 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		num = num * i
	}
	return num
}
