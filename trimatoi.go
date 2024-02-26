package piscine

func TrimAtoi(s string) int {
	ans := 0
	nan := true
	found := false
	for index, num := range s {
		if num == '-' && !found {
			for i := 0; i < index; i++ {
				if s[i] >= '0' && s[i] <= '9' {
					nan = false
				}
			}
			if nan {
				found = true
			}

		}
		if num >= '0' && num <= '9' {
			ans = ans*10 + int(num-'0')
		}

	}

	if found {
		ans *= -1
	}

	return ans
}
