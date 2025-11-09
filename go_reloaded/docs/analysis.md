# go-reloaded — Analysis

## Goal (ELI5)
Turn a messy text into a clean text by following rules exactly: Input file → Tool → Output file.

## Rules (from subject)
1) Numbers:
   - `<HEX> (hex)` → decimal (replace previous word)
   - `<BIN> (bin)` → decimal (replace previous word)
2) Casing:
   - `(up)`, `(low)`, `(cap)` affect the previous word
   - `(up, N)`, `(low, N)`, `(cap, N)` affect the previous N words (skip non-words)
3) Punctuation spacing:
   - `. , ! ? : ;` stick to the left word; a space follows them
   - Groups `...` and `!?` stay grouped; stick left; a space follows them
4) Quotes:
   - Single quotes `'` appear in pairs and hug the inner content: `'awesome'`, `'many words'`
5) Article fix:
   - `a` → `an` if the NEXT token is a word starting with vowel or `h` (case-insensitive)

## Edge cases
- Malformed markers like `(up, x)` → ignore safely
- Marker at start with no previous word → ignore
- Invalid hex/bin number → ignore (leave as-is)
- Odd number of quotes → leave as-is (don’t crash)
- Multiple `!` like `!!` → two `!` tokens; both attach left per punctuation rule

## Processing order (Pipeline)
Tokenize → Hex/Bin → Casing → Quotes → Punctuation → Article (a→an) → Smart Detokenize
