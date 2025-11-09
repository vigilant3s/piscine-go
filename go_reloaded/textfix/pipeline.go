package textfix

import "strings"

// ProcessText: main pipeline orchestrator.
func ProcessText(in string) string {
	// If the input file ends with a newline, remember it so we can keep it.
	hadNewline := strings.HasSuffix(in, "\n")

	t := Tokenize(in)   // words, markers, punctuation, quotes
	t = PassHexBin(t)   // 1E (hex) -> 30, 10 (bin) -> 2
	t = PassCasing(t)   // (up)/(low)/(cap) and (up, N)...
	t = PassQuotes(t)   // quotes as tokens; spacing handled in detokenizer
	t = PassPunct(t)    // punctuation normalization like "?." -> "?"
	t = PassArticles(t) // a -> an before vowel/h
	out := DetokenizeSmart(t)

	if hadNewline {
		out += "\n"
	}
	return out
}
