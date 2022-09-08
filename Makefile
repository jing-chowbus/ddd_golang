TAG := $(shell git describe --tags)
PROJECT_ROOT=$(shell pwd)

build:
	PROJECT_ROOT=$(PROJECT_ROOT) $(MAKE) build -C $(PROJECT_ROOT)/cmd/web/

run:
	PROJECT_ROOT=$(PROJECT_ROOT) $(MAKE) run -C $(PROJECT_ROOT)/cmd/web/
