package textfix

// PassQuotes: no token changes here; detokenizer handles hugging/spaces.
// Keep as a separate pass for clarity and future tweaks.
func PassQuotes(t []string) []string { return t }
