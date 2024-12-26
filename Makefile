# Makefile for building a single Go binary and creating symlinks.

BINARY_NAME := go-lzw
SOURCES := $(wildcard *.go)

.PHONY: all build links clean

all: build links

# Build the Go binary
build:
	go build -o $(BINARY_NAME) $(SOURCES)

# Create symlinks
links:
	# Remove existing symlinks if they exist
	rm -f encode decode
	# Create new symlinks to the compiled binary
	ln -s $(BINARY_NAME) encode
	ln -s $(BINARY_NAME) decode

# Clean up binaries and symlinks
clean:
	rm -f $(BINARY_NAME) encode decode

