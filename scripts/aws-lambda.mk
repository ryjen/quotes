export GO ?= go
export AWS_ROLE ?= arn:aws:iam::872755025701:role/Lambdadmin

FUNCTIONS_DIR ?= functions

BUILDERS += functions
HELPS += aws-functions-help
CLEANERS += aws-lambda-clean

FUNCTION_BIN = main
FUNCTION_PKG = function.zip

SUBDIRS := $(patsubst $(FUNCTIONS_DIR)/%/., %, $(wildcard $(FUNCTIONS_DIR)/*/.))
BUILD_FUNCTIONS := $(patsubst %, %.build, $(SUBDIRS))
DEPLOY_FUNCTIONS := $(patsubst %, %.deploy, $(SUBDIRS))

.PHONY: functions deploy undeploy $(SUBDIRS) aws-lambda-help aws-lambda-build aws-lambda-clean aws-lambda-package aws-lambda-deploy

functions: $(BUILD_FUNCTIONS)
%.build:
	@$(MAKE) -f $(PWD)/scripts/aws-lambda.mk -E "AWS_LAMBDA_NAME=$*" -C $(FUNCTIONS_DIR)/$* aws-lambda-build

deploy: $(DEPLOY_FUNCTIONS)
%.deploy:
	@$(MAKE) -f $(PWD)/scripts/aws-lambda.mk -E "AWS_LAMBDA_NAME=$*" -C $(FUNCTIONS_DIR)/$* aws-lambda-deploy 

undeploy: aws-lambda-undeploy

aws-functions-help:
	@echo " deploy              deploy aws lambda functions"
	@echo " undeploy            undeploy aws lambda functions"
	@echo " functions           build aws lambda functions"

aws-lambda-clean:
	@echo "cleaning function $(AWS_LAMBDA_NAME)"
	@rm -f main $(FUNCTION_PKG)

aws-lambda-build:
	@echo "building function $(AWS_LAMBDA_NAME)"
	@GOOS=linux CGO_ENABLED=0 $(GO) build -o $(FUNCTION_BIN) ./...

$(FUNCTION_BIN): aws-lambda-build
aws-lambda-package: $(FUNCTION_BIN)
	@echo "packaging function $(AWS_LAMBDA_NAME)"
	@zip $(FUNCTION_PKG) $(FUNCTION_BIN) >/dev/null 2>&1

aws-lambda-undeploy:
	@if [ -e .deployed ]; then \
	  aws lambda delete-function --function-name $(AWS_LAMBDA_NAME) >/dev/null 2>&1; \
		rm -f .deployed; \
	fi

aws-lambda-deploy: aws-lambda-package aws-lambda-undeploy
	@echo "creating function $(AWS_LAMBDA_NAME)"
	@aws lambda create-function --function-name $(AWS_LAMBDA_NAME) --zip-file fileb://$(FUNCTION_PKG) --handler main --runtime go1.x --role $(AWS_ROLE) >/dev/null
	@touch .deployed

