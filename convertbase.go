package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string { // ConvertBase converts a number from one base to another
	// Step 1: Create a map of characters to values from baseFrom
	fromMap := make(map[rune]int) // Map to store the value of each character in baseFrom
	for i, r := range baseFrom {  // Iterate over baseFrom and assign values
		fromMap[r] = i // Assign the index as the value for the character
	}

	// Step 2: Convert nbr from baseFrom to base 10
	base10 := 0                  // Initialize base10 to 0
	baseFromLen := len(baseFrom) // Length of baseFrom
	for _, r := range nbr {      // Iterate over each character in nbr
		digit := fromMap[r]                 // Get the value of the character from the map
		base10 = base10*baseFromLen + digit // Update base10
	}

	// Step 3: Convert base10 to baseTo
	if base10 == 0 { // Handle the case where base10 is 0
		return string(baseTo[0]) // Handle "0"
	}

	baseToRunes := []rune(baseTo) // Convert baseTo to runes for easier manipulation
	baseToLen := len(baseTo)      // Length of baseTo
	var result []rune             // Slice to store the result

	for base10 > 0 { // While base10 is greater than 0
		remainder := base10 % baseToLen                            // Calculate the remainder
		result = append([]rune{baseToRunes[remainder]}, result...) // Prepend the character corresponding to the remainder to the result
		base10 = base10 / baseToLen                                // Update base10
	}

	return string(result) // Convert the result slice of runes to a string and return it
}
