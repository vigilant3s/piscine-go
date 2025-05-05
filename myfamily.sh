#!/bin/bash

curl -sL -H "User-Agent: Mozilla/5.0" https://platform.zone01.gr/assets/superhero/all.json | \
jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber)) | .connections.relatives'