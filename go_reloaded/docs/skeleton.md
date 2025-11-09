# Skeleton Plan

## Files
- `main.go` — CLI: read input → ProcessText → write output
- `textfix/`
  - `tokens.go` — Tokenize(), DetokenizeSmart()
  - `hexbin.go` — PassHexBin()
  - `casing.go` — PassCasing()
  - `quotes.go` — PassQuotes()
  - `punct.go` — PassPunct()
  - `article.go` — PassArticles()
  - `pipeline.go` — ProcessText()

## TDD Tasks (small steps)
1) Tokenize tests → Tokenize
2) Smart detokenize (spacing rules)
3) Hex/Bin tests → PassHexBin
4) Casing tests → PassCasing (single + numbered)
5) Quotes tests → PassQuotes (spacing in detokenizer)
6) Punctuation tests → PassPunct + spacing rules
7) Article tests → PassArticles
8) Golden E2E: run examples to verify exact output