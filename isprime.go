package piscine

func IsPrime(nb int) bool { // Check if a number is prime
	if nb <= 1 { // Check if the number is less than or equal to 1
		return false // 1 is not prime
	}
	if nb == 2 || nb == 3 { // 2 and 3 are prime numbers
		return true // Return true
	}
	if nb%2 == 0 || nb%3 == 0 { // Check if the number is divisible by 2 or 3
		return false // Not prime
	}

	// Check from 5 to sqrt(nb) using 6-step
	for i := 5; i*i <= nb; i += 6 { // Check for factors of the form 6k ± 1
		if nb%i == 0 || nb%(i+2) == 0 { // Check if the number is divisible by i or i+2
			return false //	 Not prime
		}
	}
	return true // Return true if no factors were found
}
