#!/bin/bash

API_ENDPOINT="http://localhost:8080/albums"

read -p "Enter album id: " id

curl -X GET "$API_ENDPOINT/$id" \
  -H "Content-Type: application/json" \
