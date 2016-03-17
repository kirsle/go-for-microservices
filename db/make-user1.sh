#!/bin/bash
curl -X POST -H "Content-Type: application/json" \
	-d '{"username": "bob", "firstName": "Bob", "lastName": "Dole"}' \
	http://localhost:8000/users
