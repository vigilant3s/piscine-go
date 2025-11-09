package textfix

import (
	"strings"
	"unicode"
)

// Tokenize: turns raw string into tokens (words, markers, quotes, punctuation).
func Tokenize(s string) []string {
	var tokens []string
	runes := []rune(s)
	i := 0

	emitWord := func(start, end int) {
		if end > start {
			tokens = append(tokens, string(runes[start:end]))
		}
	}

	for i < len(runes) {
		r := runes[i]

		// Group: "..." or "!?"
		if r == '.' {
			if i+2 < len(runes) && runes[i+1] == '.' && runes[i+2] == '.' {
				tokens = append(tokens, "...")
				i += 3
				continue
			}
			tokens = append(tokens, ".")
			i++
			continue
		}
		if r == '!' && i+1 < len(runes) && runes[i+1] == '?' {
			tokens = append(tokens, "!?")
			i += 2
			continue
		}

		// Single punctuation: , ! ? : ;
		if r == ',' || r == '!' || r == '?' || r == ':' || r == ';' {
			tokens = append(tokens, string(r))
			i++
			continue
		}

		// Single quote as its own token
		if r == '\'' {
			tokens = append(tokens, "'")
			i++
			continue
		}

		// Marker: ( ... )
		if r == '(' {
			start := i
			for i < len(runes) && runes[i] != ')' {
				i++
			}
			if i < len(runes) && runes[i] == ')' {
				i++
				tokens = append(tokens, string(runes[start:i]))
				continue
			}
			// no closing ) — be robust
			tokens = append(tokens, "(")
			i++
			continue
		}

		// Space: skip
		if unicode.IsSpace(r) {
			i++
			continue
		}

		// Word: letters/digits
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			start := i
			i++
			for i < len(runes) && (unicode.IsLetter(runes[i]) || unicode.IsDigit(runes[i])) {
				i++
			}
			emitWord(start, i)
			continue
		}

		// Any other char: single token
		tokens = append(tokens, string(r))
		i++
	}

	return tokens
}

// DetokenizeSmart: renders tokens with correct spacing rules.
func DetokenizeSmart(tokens []string) string {
	var b strings.Builder

	isPunct := func(t string) bool {
		switch t {
		case ".", ",", "!", "?", ":", ";", "...", "!?":
			return true
		default:
			return false
		}
	}
	isQuote := func(t string) bool { return t == "'" }

	inQuote := false // false -> next quote is opening; true -> next quote is closing

	for i := 0; i < len(tokens); i++ {
		cur := tokens[i]
		var prev string
		if i > 0 {
			prev = tokens[i-1]
		}

		needSpace := true
		if i == 0 {
			needSpace = false
		} else {
			switch {
			case isPunct(cur):
				// Never a space BEFORE punctuation
				needSpace = false

			case isQuote(cur) && !inQuote:
				// Opening quote: usually needs a space before (e.g., after a word or colon)
				needSpace = true

			case isQuote(cur) && inQuote:
				// Closing quote: no space before it
				needSpace = false

			case isQuote(prev):
				// After an opening quote: no space before the next token
				// (This also covers the token right after we printed an opening quote)
				needSpace = false
			}
		}

		if needSpace {
			b.WriteByte(' ')
		}
		b.WriteString(cur)

		// Toggle when we see a quote
		if isQuote(cur) {
			inQuote = !inQuote
		}
	}

	return b.String()
}
