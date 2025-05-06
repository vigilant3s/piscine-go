#!/bin/bash

# Fetch the superhero data and process it
curl -s https://platform.zone01.gr/assets/superhero/all.json | \
jq -r --arg id "$HERO_ID" '
  .[] | 
  select(.id == ($id | tonumber)) | 
  .connections.relatives | 
  gsub("\""; "") | 
  gsub("\n"; " ") |  # Replace newlines with spaces
  gsub(" +"; " ")     # Replace multiple spaces with a single space
' | \
tr '\n' ';' | \
sed 's/;$/\n/'