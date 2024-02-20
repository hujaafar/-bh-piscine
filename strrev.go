package piscine

func StrRev(s string) string {
	var outpot string
	for i := len(s) - 1; i >= 0; i-- {
		outpot += string(s[i])
	}
}
