# Makefile that builds go-hello-world, a "go" program.

# "Simple expanded" variables (':=')

# PROGRAM_NAME is the name of the GIT repository.
PROGRAM_NAME := $(shell basename `git rev-parse --show-toplevel`)
MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MAKEFILE_DIRECTORY := $(dir $(MAKEFILE_PATH))
TARGET_DIRECTORY := $(MAKEFILE_DIRECTORY)/target
DOCKER_CONTAINER_NAME := $(PROGRAM_NAME)
DOCKER_IMAGE_NAME := senzing/$(PROGRAM_NAME)
DOCKER_BUILD_IMAGE_NAME := $(DOCKER_IMAGE_NAME)-build
BUILD_VERSION := $(shell git describe --always --tags --abbrev=0 --dirty)
BUILD_TAG := $(shell git describe --always --tags --abbrev=0)
BUILD_ITERATION := $(shell git log $(BUILD_TAG)..HEAD --oneline | wc -l | sed 's/^ *//')
GIT_REMOTE_URL := $(shell git config --get remote.origin.url)
GO_PACKAGE_NAME := $(shell echo $(GIT_REMOTE_URL) | sed -e 's|^git@github.com:|github.com/|' -e 's|\.git$$||')

# Recursive assignment ('=')

CC = gcc
# Conditional assignment. ('?=')

SENZING_G2_DIR ?= /opt/senzing/g2
SENZING_DATABASE_URL ?= postgresql://postgres:postgres@127.0.0.1:5432/G2

# The first "make" target runs as default.

.PHONY: default
default: help

# -----------------------------------------------------------------------------
# Export environment variables.
# -----------------------------------------------------------------------------

.EXPORT_ALL_VARIABLES:

# Flags for the C compiler

LD_LIBRARY_PATH = ${SENZING_G2_DIR}/lib

# -----------------------------------------------------------------------------
# Build
# -----------------------------------------------------------------------------

.PHONY: dependencies
dependencies:
	@go get -u ./...
	@go get -t -u ./...
	@go mod tidy

# -----------------------------------------------------------------------------
# Test
# -----------------------------------------------------------------------------

.PHONY: test
test:
#	@go test -v ./...
#	@go test -v ./.
#	@go test -v ./g2config
#	@go test -v ./g2configmgr
#	@go test -v ./g2diagnostic
#	@go test -v ./g2engine
	@go test -v ./g2product


# -----------------------------------------------------------------------------
# Run
# -----------------------------------------------------------------------------

.PHONY: run
run:
	@go run main.go

# -----------------------------------------------------------------------------
# Utility targets
# -----------------------------------------------------------------------------

.PHONY: clean
clean:
	@go clean -cache
	@docker rm --force $(DOCKER_CONTAINER_NAME) 2> /dev/null || true
	@docker rmi --force $(DOCKER_IMAGE_NAME) $(DOCKER_BUILD_IMAGE_NAME) 2> /dev/null || true
	@rm -rf $(TARGET_DIRECTORY) || true
	@rm -f $(GOPATH)/bin/$(PROGRAM_NAME) || true


.PHONY: print-make-variables
print-make-variables:
	@$(foreach V,$(sort $(.VARIABLES)), \
		$(if $(filter-out environment% default automatic, \
		$(origin $V)),$(warning $V=$($V) ($(value $V)))))


.PHONY: help
help:
	@echo "Build $(PROGRAM_NAME) version $(BUILD_VERSION)-$(BUILD_ITERATION)".
	@echo "All targets:"
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$' | xargs
