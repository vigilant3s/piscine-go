package textfix

import (
	"strconv"
	"strings"
	"unicode"
)

func PassCasing(t []string) []string {
	out := make([]string, 0, len(t))

	isWord := func(s string) bool {
		if s == "" {
			return false
		}
		for _, r := range s {
			if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
				return false
			}
		}
		return true
	}

	applyCap := func(s string) string {
		if s == "" {
			return s
		}
		rs := []rune(strings.ToLower(s))
		rs[0] = unicode.ToUpper(rs[0])
		return string(rs)
	}

	parseMarker := func(m string) (kind string, n int, ok bool) {
		if !(strings.HasPrefix(m, "(") && strings.HasSuffix(m, ")")) {
			return "", 0, false
		}
		body := strings.TrimSuffix(strings.TrimPrefix(m, "("), ")")
		parts := strings.Split(body, ",")
		parts[0] = strings.TrimSpace(parts[0])

		switch parts[0] {
		case "up", "low", "cap":
			kind = parts[0]
		default:
			return "", 0, false
		}

		if len(parts) == 1 {
			return kind, 1, true
		}
		if len(parts) == 2 {
			clean := strings.TrimSpace(parts[1])
			val, err := strconv.Atoi(clean)
			if err != nil || val < 0 {
				return "", 0, false
			}
			return kind, val, true
		}
		return "", 0, false
	}

	for i := 0; i < len(t); i++ {
		cur := t[i]
		kind, n, ok := parseMarker(cur)
		if !ok {
			out = append(out, cur)
			continue
		}

		// Apply to previous n WORD tokens (skip punctuation/markers).
		count := 0
		for j := len(out) - 1; j >= 0 && count < n; j-- {
			if !isWord(out[j]) {
				continue
			}
			switch kind {
			case "up":
				out[j] = strings.ToUpper(out[j])
			case "low":
				out[j] = strings.ToLower(out[j])
			case "cap":
				out[j] = applyCap(out[j])
			}
			count++
		}
		// do not append the marker itself
	}

	return out
}
