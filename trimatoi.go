package piscine

func TrimAtoi(s string) int { // convert string to int
	num := 0            // initialize num to 0
	foundDigit := false // flag to check if we found a digit
	negative := false   // flag to check if the number is negative

	for _, ch := range s { // iterate over each character in the string
		if ch == '-' && !foundDigit { // check for negative sign
			negative = true // set negative flag if we haven't found a digit yet
		} // check for negative sign
		if ch >= '0' && ch <= '9' { // check if the character is a digit
			foundDigit = true          // set foundDigit flag to true
			num = num*10 + int(ch-'0') // build the number digit by digit
		} // check if the character is a digit
	} // check for digit

	if !foundDigit { // if no digits were found, return 0
		return 0
	}

	if negative { // if the number is negative, return its negative value
		return -num
	}

	return num // return the number
}
