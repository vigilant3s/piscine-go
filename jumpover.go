package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	result := []rune{}
	for i, r := range str {
		if (i+1)%3 == 0 {
			result = append(result, r)
		}
	}

	return string(result) + "\n"
}
