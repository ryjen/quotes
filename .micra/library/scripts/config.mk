HELPS ?=
BUILDERS ?=
CLEANERS ?=
INSTALLERS ?=

CMD ?= cmd/$(PROJECT_NAME)
EXE ?= $(PROJECT_NAME)

BUILD_FLAGS ?=
LINK_FLAGS ?=

FEATURE_ENABLED := 0
FEATURE_CMD ?= $(shell test -e $(CMD); echo $$?)
FEATURE_WEB ?= $(shell test -d web; echo $$?)
FEATURE_PLUGIN ?= $(shell test -d plugin; echo $$?)
