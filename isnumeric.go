package piscine

func IsNumeric(nbr int) bool {
	for i := 0; i < nbr; i++ {
		if nbr < 48 && nbr > 57 {
			return false
		}
	}
	return true
}
