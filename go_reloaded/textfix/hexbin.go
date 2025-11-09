package textfix

import (
	"strconv"
	"strings"
)

// PassHexBin: convert "NN (hex)" or "NN (bin)" to decimal (replace NN, drop marker).
func PassHexBin(t []string) []string {
	out := make([]string, 0, len(t))
	for i := 0; i < len(t); i++ {
		cur := t[i]
		switch cur {
		case "(hex)":
			if len(out) > 0 {
				prev := out[len(out)-1]
				if dec, err := strconv.ParseInt(prev, 16, 64); err == nil {
					out[len(out)-1] = strconv.FormatInt(dec, 10)
				}
			}
			// drop marker
		case "(bin)":
			if len(out) > 0 {
				prev := out[len(out)-1]
				if dec, err := strconv.ParseInt(prev, 2, 64); err == nil {
					out[len(out)-1] = strconv.FormatInt(dec, 10)
				}
			}
			// drop marker
		default:
			out = append(out, cur)
		}
	}
	return out
}

// eqCI: case-insensitive equals
func eqCI(a, b string) bool {
	return strings.EqualFold(a, b)
}
