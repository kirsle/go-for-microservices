#!/bin/bash
curl -X POST -H "Content-Type: application/json" \
	-d '{"username": "alice", "firstName": "Alice", "lastName": "Alison"}' \
	http://localhost:8000/users
