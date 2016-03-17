#!/bin/bash

curl -v -H "Content-Type: application/json" \
	-u foo:bar \
	-d '{"value": "Hello world"}' \
	-X POST \
	http://localhost:8000/user
