package piscine

func IsPrime(nbr int) bool {
	if nbr < 2 {
		return false
	}
	for i := 2; i*i <= nbr; i++ {
		if nbr%i == 0 {
			return false
		}
	}
	return true
}
