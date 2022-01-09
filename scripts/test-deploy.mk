IllegalLocationConstraintException
RANDOM := $(shell bash -c 'echo $$RANDOM')

test-deploy:
	curl \
		-d "{ \"text\": \"This is quote #$(RANDOM).\", \"author\": \"$(USER)\" }" \
		-H 'Content-Type: application/json' \
		https://ikfi3oi9te.execute-api.ca-central-1.amazonaws.com/staging/quotes 
	curl \
		-H 'Accept: application/json' \
		https://ikfi3oi9te.execute-api.ca-central-1.amazonaws.com/staging/quotes 