.PHONY: web-dist web-dev web-setup

DISTBUILDS += web-dist

CLEANERS += web-clean

ifneq ($(FEATURE_CMD),$(FEATURE_ENABLED))

DEVBUILDS += web-dev

else

HELPS += web-help

.PHONY: web
web: web-dev

.PHONY: web-help
web-help:
	@echo " web                 build the dev web frontend"

endif

web-dist: web-setup
	@yarn --mode=production --cwd web build

web-dev: web-setup
	@yarn --mode=development --cwd web build:dev

web-setup:
	@yarn --cwd web install

web-clean: web-setup
	@yarn --cwd web clean

