STACK := quotes

FUNCTIONS_DIR := cmd
OUTPUT_DIR := bin

export GOOS ?= linux
export GOARCH ?= amd64
export CGO_ENABLED ?= 0

FUNCTIONS := $(patsubst $(FUNCTIONS_DIR)/%/., %, $(wildcard $(FUNCTIONS_DIR)/*/.))
BUILD_FUNCTIONS := $(patsubst %, %.build, $(FUNCTIONS))

help:
	@echo "build    build lambda functions"
	@echo "test     run tests"
	@echo "api      run locally"
	@echo "deploy   deploy the functions"

.PHONY: functions
functions: $(OUTPUT_DIR) $(BUILD_FUNCTIONS)

$(OUTPUT_DIR):
	@mkdir -p $@

.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: test
test:
	@echo "1. Test create single quote"
	sam local invoke CreateQuote -e test/create-quote.json 2>/dev/null
	@echo -e "\n\n2. Test create bulk quotes"
	sam local invoke CreateQuote -e test/create-quotes.json 2>/dev/null
	@echo -e "\n\n3. Test read single quote"
	sam local invoke ReadQuote -e test/read-quote.json 2>/dev/null
	@echo -e "\n\n4. Test list all quotes"
	sam local invoke ReadQuote -e test/read-quotes.json 2>/dev/null
	@echo -e "\n\n5. Test page quotes"
	sam local invoke ReadQuote -e test/page-quotes.json 2>/dev/null
	@echo -e "\n\n6. Test update quote"
	sam local invoke UpdateQuote -e test/update-quote.json 2>/dev/null
	@echo -e "\n\n7. Test delete quote"
	sam local invoke DeleteQuote -e test/delete-quote.json 2>/dev/null
	# TODO: parse id

.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: install
install:
	go get -d ./...
	go get -t ./...

%.build: 
	go build -o $(OUTPUT_DIR)/$* ./$(FUNCTIONS_DIR)/$*

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda: functions

.PHONY: build
build: clean unit-test lambda
	sam build

.PHONY: api
api: build
	sam local start-api

.PHONY: deploy
deploy: build
	sam deploy --resolve-s3 --stack-name $(STACK) --capabilities CAPABILITY_IAM
