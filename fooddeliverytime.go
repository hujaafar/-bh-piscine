package piscine

type food struct {
	preptime int
}

var menu = map[string]food{

	"burger":  {preptime: 15},
	"chips":   {preptime: 10},
	"nuggets": {preptime: 12},
}

func FoodDeliveryTime(order string) int {
	value, error := menu[order]
	if !error {
		return 0
	}
	return value.preptime
}
