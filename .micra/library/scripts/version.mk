MICRA := $(shell command -v micra 2>/dev/null)

.PHONY: get-version
get-version:
	$(if $(MICRA),$(eval BUILD_VERSION ?= $(shell $(MICRA) build version $(VERSION_FLAGS) 2>/dev/null)),$(eval BUILD_VERSION ?= "v0.0.1-alpha+1"))

.PHONY: link-version
link-version: get-version
	$(if $(BUILD_VERSION),$(eval LINK_FLAGS += -X main.version=$(BUILD_VERSION)))

.PHONY: increment-version
increment-version:
	$(eval VERSION_FLAGS := --inc)

.PHONY: build-version
build-version: increment-version link-version

.PHONY: install-version
install-version: link-version

.PHONY: help-version
help-version: get-version
	@echo ""
ifdef BUILD_VERSION
	@echo "version: $(BUILD_VERSION)"
endif
