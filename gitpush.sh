#!/bin/bash

# If no filename is provided, push everything with default message
if [ -z "$1" ]; then
  echo "No filename provided. Adding and committing all changes..."
  git add .
  git commit -m "Auto-commit: push all changes"
  git push
  exit 0
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

# Git operations for single file
git add "$FILENAME"
git commit -m "$COMMIT_MSG"
git push
