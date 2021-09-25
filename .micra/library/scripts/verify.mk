
GOLINT ?= golint

TEST_TAGS ?= test

HELPS += help-format help-lint help-test

VERIFY_PKGS := $(shell go list ./...)

$(GOLINT):
	@$(GO) get -u golang.org/x/lint/golint

.PHONY: format
format:
	@echo "Formatting..."
	@$(GO) fmt $(VERIFY_PKGS)

.PHONY: help-format
help-format:
	@echo " format              format source code"

.PHONY: lint
lint: format $(GOLINT)
	@echo "Linting..."
	@$(GOLINT) $(VERIFY_PKGS)

.PHONY: help-lint
help-lint:
	@echo " lint                run source code linter"

.PHONY: test
test:
	@echo "Testing..."
	@$(GO) test -tags $(TEST_TAGS) $(VERIFY_PKGS)

.PHONY: verify
verify: lint test

.PHONY: help-test
help-test:
	@echo " test                test source code"
	@echo " verify              lint and test source code"
