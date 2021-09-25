
PLUGIN_NAMES := $(shell ls -d plugin/*/ | cut -d "/" -f 2)

INSTALL_PATH := $(MICRA_PATH)/library/plugin

PLUGIN_MODULES := $(patsubst %, $(BIN_DIR)/%.so, $(PLUGIN_NAMES))

INSTALL_PLUGINS := $(patsubst %, %.install, $(PLUGIN_NAMES))

CLEAN_PLUGINS := $(patsubst %, %.clean, $(PLUGIN_NAMES))

HELPS += help-plugin

BUILDERS += plugin

INSTALLERS += install-plugin

CLEANERS += clean-plugin

.PHONY: help-plugin
help-plugin:
	@echo " plugin              make plugins"

## Building

$(BIN_DIR)/%.so:
	@echo "Building $* plugin..."
	@$(GO) generate $(BUILD_FLAGS) ./plugin/$*
	@$(GO) build $(BUILD_FLAGS) -ldflags "$(LINK_FLAGS)" -buildmode=plugin -o $@ ./plugin/$*/generate.go

.PHONY: plugin
plugin: $(PLUGIN_MODULES)

## Cleaning

.PHONY: clean-plugin
clean-plugin: $(CLEAN_PLUGINS)

%.clean:
	@echo "Cleaning $* plugin"
	@rm -f $(BIN_DIR)/$*.so $(INSTALL_PATH)/$*.so

## Installing

$(INSTALL_PATH):
	@mkdir -p $@

%.install:
	@echo "Installing $*.so"
	@install $(BIN_DIR)/$*.so $(INSTALL_PATH)/$*.so

.PHONY: install-plugin
install-plugin: $(INSTALL_PATH) $(INSTALL_PLUGINS)

