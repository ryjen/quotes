
test-deploy:
	curl \
		-X POST \
		-d '{ "text": "This is a quote.", "author": "$(USER)" }' \
		-H 'Content-Type: application/json' \
		https://ikfi3oi9te.execute-api.ca-central-1.amazonaws.com/staging/quotes 