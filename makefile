# Define Variables
OUTPUT_DIRS := bin/* pkg/*
GOCOMMAND := go
BUILD_PKGS := mathutil geometry-app json-parsing flags file-io

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
