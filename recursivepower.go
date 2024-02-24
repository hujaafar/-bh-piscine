package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	if nb <= 20 && nb >= 0 {
		return nb * RecursivePower(nb, power-1)
	} else {
		return 0
	}
}
