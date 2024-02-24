package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 { // Negative values are not allowed
		return 0
	}

	if nb == 0 {
		return 1
	}

	num := 1
	for i := 1; i <= nb; i++ {
		// Check for potential overflow
		if num > (1<<31-1)/i {
			return 0
		}
		num = num * i
	}

	return num
}
