#!/bin/bash

# Check if filename is provided
if [ -z "$1" ]; then
  echo "Usage: ./gitpush.sh filename [optional commit message]"
  exit 1
fi

FILENAME="$1"
COMMIT_MSG="${2:-Auto-commit for $FILENAME}"

# Check if file exists
if [ ! -f "$FILENAME" ]; then
  echo "Error: File '$FILENAME' not found."
  exit 1
fi

# Make executable if it's a shell script
if [[ "$FILENAME" == *.sh ]]; then
  chmod +x "$FILENAME"
fi

# Git operations
git add "$FILENAME"
git commit -m "$COMMIT_MSG"
git push