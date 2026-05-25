#!/bin/bash

URI="https://acad.learn2earn.ng/assets/superhero/all.json"

curl -s $URI | jq -r --argjson ID "$HERO_ID" ' .[] | select(.id == $ID) | .name, .powerstats.power,.appearance.gender '