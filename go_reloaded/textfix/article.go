package textfix

import "unicode"

func PassArticles(t []string) []string {
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
	startsWithVowelOrH := func(s string) bool {
		if s == "" {
			return false
		}
		r := unicode.ToLower([]rune(s)[0])
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'h':
			return true
		default:
			return false
		}
	}

	for i := 0; i < len(t); i++ {
		if t[i] == "a" || t[i] == "A" {
			// look ahead to the NEXT token; only convert if it's a WORD
			if i+1 < len(t) && isWord(t[i+1]) && startsWithVowelOrH(t[i+1]) {
				if t[i] == "a" {
					t[i] = "an"
				} else {
					t[i] = "An"
				}
			}
		}
	}
	return t
}
