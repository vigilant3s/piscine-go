package textfix

// PassPunct normalizes tricky punctuation sequences before spacing is applied.
// Rules:
//  1. If a '?' or '!' is immediately followed by '.', drop the '.' (e.g., "?.", "!." -> "?" / "!").
//  2. If we see PATTERN: <end-punct> "'" <same end-punct>, drop the OUTSIDE one,
//     keeping punctuation inside the quotes. Example: ".'." -> ".'", "?'?" -> "?'".
func PassPunct(t []string) []string {
	out := make([]string, 0, len(t))

	isEnd := func(s string) bool {
		return s == "." || s == "!" || s == "?"
	}

	for i := 0; i < len(t); i++ {
		cur := t[i]

		// Rule 1: ?/! followed by .  => drop the .
		if (cur == "?" || cur == "!") && i+1 < len(t) && t[i+1] == "." {
			out = append(out, cur)
			i++ // skip the following '.'
			continue
		}

		// Rule 2: <end> "'" <same end>  => keep inside one, drop outside
		if isEnd(cur) && i+2 < len(t) && t[i+1] == "'" && t[i+2] == cur {
			// emit end-punct and the quote; skip the outside duplicate end-punct
			out = append(out, cur)
			out = append(out, "'")
			i += 2
			continue
		}

		out = append(out, cur)
	}
	return out
}
