#!/bin/bash

curl -v -H "Content-Type: application/json" \
	-u admin:passwd \
	-d '{"value": "Basic API with Gin Framework"}' \
	-X POST \
	http://localhost:8000/user
