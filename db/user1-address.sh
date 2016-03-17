#!/bin/bash
curl -X POST -H "Content-Type: application/json" \
	-d '{"street1": "8520 National Blvd", "city": "Culver City", "state": "CA", "zip": "90232"}' \
	http://localhost:8000/user/1/address
