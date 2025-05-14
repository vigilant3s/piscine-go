package piscine

func IsAlpha(s string) bool { // IsAlpha checks if a string contains only alphabetic characters (a-z, A-Z) and digits (0-9).
	for _, r := range s { // iterate over each character in the string
		if (r < 'a' || r > 'z') && // check for lowercase letters
			(r < 'A' || r > 'Z') && // check for uppercase letters
			(r < '0' || r > '9') { // check for digits
			return false // if any character is not in the valid ranges, return false
		}
	}
	return true // if all characters are valid, return true
}

// IsAlpha checks if a string contains only alphabetic characters (a-z, A-Z) and digits (0-9).
// It returns true if the string is valid, otherwise false.
// The function iterates over each character in the string and checks if it falls within the valid ranges.
// If any character is outside these ranges, the function returns false.
// If all characters are valid, it returns true.
// The function uses a range loop to iterate over the string, which allows it to handle Unicode characters correctly.
// The function is case-sensitive, meaning it treats uppercase and lowercase letters as valid characters.
// The function does not modify the input string or create any new strings, making it efficient in terms of memory usage.
// The function is simple and efficient, with a time complexity of O(n), where n is the length of the input string.
// It is a good practice to use this function when validating user input or processing strings that should only contain alphabetic characters and digits.
