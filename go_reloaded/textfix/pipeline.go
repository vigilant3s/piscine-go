package textfix

// ProcessText: main pipeline orchestrator.
func ProcessText(in string) string {
	t := Tokenize(in)   // words, markers, punctuation, quotes
	t = PassHexBin(t)   // 1E (hex) -> 30, 10 (bin) -> 2
	t = PassCasing(t)   // (up)/(low)/(cap) and (up, N)...
	t = PassQuotes(t)   // leave quotes as tokens (spacing done in detokenizer)
	t = PassPunct(t)    // keep as pass; spacing handled by detokenizer
	t = PassArticles(t) // a -> an before vowel/h
	out := DetokenizeSmart(t)
	return out
}
