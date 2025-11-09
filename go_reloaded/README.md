# 🌀 go-reloaded

A small **CLI tool written in Go** that reads a text file, applies a set of transformation rules, and writes the corrected result to an output file.

This project focuses on **string manipulation**, **clean architecture**, and **test-driven development**.

---

## 🧠 Overview

`go-reloaded` reads an input text file, applies multiple logical text corrections, and produces an exact, formatted output file.

The tool respects standard Go conventions (`gofmt`, `goimports`) and uses only the **standard library**.

---

## 🧩 Features

### 🔢 Number Conversion
- `NN (hex)` → converts the hexadecimal number to decimal.  
- `NN (bin)` → converts the binary number to decimal.  
  Example:  
  `1E (hex)` → `30`, `10 (bin)` → `2`.

### ✍️ Casing
- `(up)`, `(low)`, `(cap)` affect the previous word.  
- `(up, N)`, `(low, N)`, `(cap, N)` affect the previous **N words**.  
  Example:  
  `this is great (up, 2)` → `this is GREAT`.

### 🔤 Article Rule
- Converts `a` → `an` if the **next word** begins with a vowel or `h`.  
  Example:  
  `a apple` → `an apple`.

### ✨ Punctuation Spacing
- `. , ! ? : ;` stick to the **previous word** and are followed by one space.  
- Punctuation groups like `...` or `!?` are preserved correctly.  
  Example:  
  `Hello ,world !!` → `Hello, world!!`

### 🗣️ Quotes
- Single quotes `' '` come in pairs and must **hug** the enclosed text.  
  Example:  
  `" ' I am happy ' "` → `'I am happy'`.

---

## ⚙️ How to Run

From the project directory:


### Navigate to project
- cd go_reloaded

### Example: convert sample text
- echo "Simply add 1E (hex) and 10 (bin)." > sample.txt

### Run the program
- go run . sample.txt result.txt

### View the result
- cat result.txt
### → Simply add 30 and 2.

## 🧱 Project Architecture
```
go_reloaded/
├── README.md
├── main.go                 # Entry point
├── textfix/                # Core transformation logic
│   ├── tokens.go           # Tokenization (splitting words/symbols)
│   ├── hexbin.go           # (hex), (bin) conversions
│   ├── casing.go           # (up), (low), (cap)
│   ├── article.go          # 'a' → 'an' logic
│   ├── punct.go            # punctuation spacing and grouping
│   ├── quotes.go           # single quote handling
│   ├── pipeline.go         # runs transformations in sequence
│   ├── e2e_test.go         # full golden tests (audit cases)
│   └── hexbin_test.go      # unit test for number conversions
├── testdata/               # Input/output reference files
│   ├── audit1.in / .out
│   ├── audit2.in / .out
│   ├── audit3.in / .out
│   └── audit4.in / .out
├── docs/
│   ├── analysis.md         # transformation rules, edge cases
│   └── skeleton.md         # step-by-step plan (TDD structure)
└── audit_check.sh          # helper script for validation
```
## 🧪 Audit & Validation
This project includes an optional helper script for validating functionality.

## 🔍 Purpose
The script automatically runs several example cases that test the program’s transformation rules (number conversion, casing, punctuation, quotes, and article logic).
It ensures the tool behaves exactly as specified.

## ⚙️ How to Run
From the root of the project, copy this: ./audit_check.sh

## 🧹 What It Does
- Creates temporary input/output files for each case
- Runs the program using go run . sample.txt result.txt
- Compares the result with the expected output
- Cleans up all temporary files automatically
- Displays ✅ PASS or ❌ FAIL for each test
```
✅ Example Output

🧪 Running Case 1...
✅ PASS  Case 1

🧪 Running Case 2...
✅ PASS  Case 2

🧪 Running Case 3...
✅ PASS  Case 3

🧪 Running Case 4...
✅ PASS  Case 4

----------------------------------
🎯 Audit Script Complete!
All temporary files have been cleaned up.
----------------------------------
```
This ensures the program’s behavior matches all expected audit cases consistently.

## 🧰 Development Notes
- Written entirely in Go 1.24.2
- Uses only standard packages
- Formatted with go fmt ./...
- Tested via go test ./textfix -v
- Complies with DRY, KISS, and SOC principles

## 💡 Author
- Dimitris Galanis
- Git Repository: https://platform.zone01.gr/git/dgalanis/go-reloaded