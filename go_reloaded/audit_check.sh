#!/bin/bash
# -------------------------------------
# Go Reloaded - Automated Audit Script
# -------------------------------------
# Runs all 4 functional audit tests defined by Zone01.
# Requires: bash, go installed, and main.go in the same folder.

set -e

# Colors for output
GREEN="\033[0;32m"
RED="\033[0;31m"
RESET="\033[0m"

# Helper function to test and print result
run_test() {
    local name="$1"
    local input="$2"
    local expected="$3"

    echo -e "\n🧪 Running $name..."
    echo "$input" > sample.txt
    go run . sample.txt result.txt >/dev/null 2>&1

    actual=$(cat result.txt | tr -d '\r')

    if [[ "$actual" == "$expected" ]]; then
        echo -e "${GREEN}✅ PASS${RESET}  $name"
    else
        echo -e "${RED}❌ FAIL${RESET}  $name"
        echo "Expected: '$expected'"
        echo "Got:      '$actual'"
    fi
}

# --------------- TEST CASES ---------------

run_test "Case 1" \
"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?" \
"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"

run_test "Case 2" \
"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure" \
"I have to pack 5 outfits. Packed 26 just to be sure"

run_test "Case 3" \
"Don not be sad ,because sad backwards is das . And das not good" \
"Don not be sad, because sad backwards is das. And das not good"

run_test "Case 4" \
"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '" \
"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"

echo -e "\n----------------------------------"
echo -e "🎯 Audit Script Complete!"
echo -e "Check ✅ marks above for passing cases."
echo -e "----------------------------------"
