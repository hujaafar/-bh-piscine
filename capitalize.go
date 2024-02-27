package piscine

func Alpha(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	array := []rune(s)

	lower := true

	for i := 0; i < len(s); i++ {
		if Alpha(array[i]) == true && lower {
			if array[i] >= 'a' && array[i] <= 'z' {
				array[i] -= 32
			}
			lower = false
		} else if array[i] >= 'A' && array[i] <= 'Z' {
			array[i] = 'a' - 'A' + array[i]
		} else if Alpha(array[i]) == false {
			lower = true
		}
	}
	return string(array)
}
