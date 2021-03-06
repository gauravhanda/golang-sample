# Define Variables
OUTPUT_DIRS := bin/* pkg/*
GOCOMMAND := go
RELATIVE_PATH := github.com/gauravhanda/golang-sample
BUILD_PKGS := $(RELATIVE_PATH)/mathutil $(RELATIVE_PATH)/geometry-app $(RELATIVE_PATH)/json-parsing \
			  $(RELATIVE_PATH)/flags $(RELATIVE_PATH)/file-io  $(RELATIVE_PATH)/httpserver

# Default target is to build the packages but clean them before
build	: clean
	$(GOCOMMAND) build $(BUILD_PKGS)

# Delete all packages and install files
clean	:
	rm -rf bin/* pkg/*

# Execute tests for all the packages
test	: build
		$(GOCOMMAND) test $(BUILD_PKGS)

# Finally install the packages to be used
install	: test
		$(GOCOMMAND) install $(BUILD_PKGS)
