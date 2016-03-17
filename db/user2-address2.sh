#!/bin/bash
curl -X POST -H "Content-Type: application/json" \
	-d '{"street1": "123 Fake Street", "street2": "Apt 1", "city": "Portland", "state": "OR", "zip": "97201"}' \
	http://localhost:8000/user/abc/address
