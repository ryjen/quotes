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

.PHONY: test
test:
	go test ./...
	sam local invoke CreateQuote -e test/create-quote.json
	# TODO: parse id

.PHONY: clean
clean:
	rm -rf $(OUTPUT_DIR)

.PHONY: install
install:
	go get -d ./...
	go get -t ./...

%.build: 
	go build -o $(OUTPUT_DIR)/$* $(FUNCTIONS_DIR)/$*/main.go

# compile the code to run in Lambda (local or real)
.PHONY: lambda
lambda: functions

.PHONY: build
build: clean lambda
	sam build

.PHONY: api
api: build
	sam local start-api

.PHONY: deploy
deploy: build
	sam deploy --resolve-s3 --stack-name $(STACK) --capabilities CAPABILITY_IAM
