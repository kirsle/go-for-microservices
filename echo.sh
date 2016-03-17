#!/bin/bash

curl -X POST -H "Content-Type: application/json" \
	-d '{"in": "Hello MediaTemple"}' \
	http://localhost:8000/echo
