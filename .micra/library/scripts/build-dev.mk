
HELPS += help-build-dev

.PHONY: help-build-dev
help-build-dev:
	@echo " dev                 make a development build"

ifeq ($(FEATURE_CMD),$(FEATURE_ENABLED))

DEVBUILDS += build-dev

.PHONY: build-dev-flags
build-dev-flags:
	$(eval BUILD_FLAGS += -tags dev)

.PHONY: build-dev
build-dev: build-version build-dev-flags
	@echo "Building development $(EXE) $(BUILD_VERSION)..."
	@$(GO) build $(BUILD_FLAGS) -ldflags "$(LINK_FLAGS)" -v -o $(BIN_DIR)/$(EXE) $(CMD)

endif

.PHONY: dev
dev: $(BUILDERS) $(DEVBUILDS)


