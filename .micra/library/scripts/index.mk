
# Makefile config

include $(MICRA_PATH)/library/scripts/config.mk

# versioning

include $(MICRA_PATH)/library/scripts/version.mk

# testing and linting

include $(MICRA_PATH)/library/scripts/verify.mk

# plugins

ifeq ($(FEATURE_PLUGIN),$(FEATURE_ENABLED))
include $(MICRA_PATH)/library/scripts/plugin.mk
endif

# Web App Building

ifeq ($(FEATURE_WEB),$(FEATURE_ENABLED))
include $(MICRA_PATH)/library/scripts/web.mk
endif

# custom scripts if they exist

-include scripts/*.mk

# build scripts

include $(MICRA_PATH)/library/scripts/build-dev.mk

include $(MICRA_PATH)/library/scripts/build-dist.mk

# project cleaning

include $(MICRA_PATH)/library/scripts/clean.mk

# help scripts

include $(MICRA_PATH)/library/scripts/help.mk
