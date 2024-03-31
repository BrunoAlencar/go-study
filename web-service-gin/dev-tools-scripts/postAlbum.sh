#!/bin/bash

API_ENDPOINT="http://localhost:8080/albums"

read -p "Enter album id: " id
read -p "Enter album title: " title
read -p "Enter album artist: " artist
read -p "Enter album price: " price

curl -X POST "$API_ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{
     "id": "'"$id"'",
     "title": "'"$title"'",
     "artist": "'"$artist"'",
     "price": '$price'
      }'