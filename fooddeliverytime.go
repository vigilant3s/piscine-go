package piscine

// Define a struct to store food preparation time
type food struct {
	preptime int
}

// FoodDeliveryTime returns the cooking time of an order item or 404 if not found
func FoodDeliveryTime(order string) int {
	// Create a menu map where each item is a food struct with prep time
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}

	// Check if the order exists in the menu
	if item, ok := menu[order]; ok {
		return item.preptime
	}

	// If not found, return error code 404
	return 404
}
