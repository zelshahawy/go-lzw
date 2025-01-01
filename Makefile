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
	rm -f encode decode
	ln -s $(BINARY_NAME) encode
	ln -s $(BINARY_NAME) decode

# Clean up binaries and symlinks
clean:
	rm -f $(BINARY_NAME) encode decode *.lzw *.out

