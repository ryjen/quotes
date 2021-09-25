

HELPS += help-build-dist

.PHONY: help-build-dist
help-build-dist:
	@echo " dist                make a distribution build"
ifeq ($(FEATURE_CMD),$(FEATURE_ENABLED))
	@echo " install             install a distribution build"
endif
ifneq ($(LOCAL_LIB),$(FEATURE_ENABLED))
	@echo " local               convert micra shared library to local project"
endif

ifeq ($(FEATURE_CMD),$(FEATURE_ENABLED))

VENDOR := vendor
DISTBUILDS += build-dist

$(VENDOR):
	@echo "Vendoring modules..."
	@$(GO) mod vendor

.PHONY: build-dist-flags
build-dist-flags: $(VENDOR)
	$(eval BUILD_FLAGS += -mod vendor)
	$(eval LINK_FLAGS += -s)

.PHONY: build-dist
build-dist: build-version build-dist-flags
	@echo "Building vendored $(EXE) $(BUILD_VERSION)..."
	@$(GO) build $(BUILD_FLAGS) -ldflags "$(LINK_FLAGS)" -o $(BIN_DIR)/$(EXE) $(CMD)

.PHONY: install
install: install-version build-dist-flags $(INSTALLERS)
	@echo "Installing $(EXE) $(BUILD_VERSION)..."
	@go install $(BUILD_FLAGS) -ldflags "$(LINK_FLAGS)" $(CMD)

endif

.PHONY: dist
dist: $(BUILDERS) $(DISTBUILDS)

ifneq ($(LOCAL_LIB),$(FEATURE_ENABLED))

.PHONY: local
local:
	@mkdir -p .micra 2>/dev/null
	@cp -rf $(MICRA_PATH)/library .micra 2>/dev/null
	@test ! -d $(MICRA_PATH)/project/$(PROJECT_NAME) || cp -rf $(MICRA_PATH)/project/$(PROJECT_NAME) .micra

endif
