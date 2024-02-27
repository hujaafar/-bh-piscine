package piscine

func BasicJoin(elems []string) string {
	my_elems := ""
	for _, w := range elems {
		my_elems = my_elems + w
	}
	return my_elems
}
