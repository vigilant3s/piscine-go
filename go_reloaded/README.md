# go-reloaded

A small CLI tool that reads a text file, applies a set of transformations, and writes the result.

## What it does (overview)

- **Number conversion**: `NN (hex)` → decimal, `NN (bin)` → decimal (replaces the previous word).
- **Casing**: `(up)`, `(low)`, `(cap)` affect the previous word; numbered versions like `(up, 3)` affect the previous N words.
- **Punctuation spacing**: `. , ! ? : ;` stick to the previous word; a space follows them. Groups `...` and `!?` are preserved and stick to the previous word.
- **Quotes**: single quotes come in pairs and should hug the words inside: `'awesome'`, `'many words'`.
- **Article rule**: `a` → `an` if the **next word** begins with a vowel or `h`.

## How to run

From the repository root (module is `piscine`):

```bash
# run from the go_reloaded directory
cd go_reloaded

# Example:
echo "Simply add 1E (hex) and 10 (bin)." > sample.txt
go run . sample.txt result.txt
cat result.txt
# -> Simply add 30 and 2.
