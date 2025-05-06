#!/bin/bash

# Check if filename is provided
if [ -z "$1" ]; then
  echo "Usage: ./gitpush.sh filename"
  exit 1
fi

# Use the filename and a default commit message if not given
FILENAME="$1"
COMMIT_MSG="${2:-Auto-commit for $FILENAME}"

# Make file executable (optional)
chmod +x "$FILENAME"

# Git add, commit and push
git add "$FILENAME"
git commit -m "$COMMIT_MSG"
git push