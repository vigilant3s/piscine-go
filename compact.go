package piscine

func Compact(ptr *[]string) int {
	a := *ptr
	j := 0 // index for next non-empty slot

	for _, v := range a {
		if v != "" {
			a[j] = v
			j++
		}
	}
	*ptr = a[:j] // slice now contains only non-empty strings
	return j
}
