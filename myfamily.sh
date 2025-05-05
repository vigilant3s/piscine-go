#!/bin/bash

curl -s https://platform.zone01.gr/assets/superhero/all.json | \
jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives | if type == "array" then join(", ") else . end'