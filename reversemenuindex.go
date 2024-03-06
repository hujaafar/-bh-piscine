package piscine

func ReverseMenuIndex(menu []string) []string {
	length := len(menu)
	reversedMenu := make([]string, length)

	for i := length - 1; i >= 0; i-- {
		reversedMenu[length-1-i] = menu[i]
	}

	return reversedMenu
}
