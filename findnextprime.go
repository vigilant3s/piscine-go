package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	// Keep checking next numbers until a prime is found
	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}

// Reuse the optimized IsPrime function
func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 || nb == 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	for i := 5; i*i <= nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}
