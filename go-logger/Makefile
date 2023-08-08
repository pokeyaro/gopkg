# Makefile for go-logger
# Author: Pokeya
# Version: 1.0

# Set phony targets
.PHONY: merge help test html

# Define variable
LOCAL_PACKAGE := github.com/pokeyaro/gopkg/go-logger

# Default target: print available commands
.PHONY: help
.DEFAULT_GOAL := help
help:
	@echo "Usage:"
	@echo "  make merge         Merge dev branch into master branch"
	@echo "  make test          Run tests and perform code checks"
	@echo "  make html          Generate HTML coverage report"

# Merge target: merge dev branch into master branch
merge:
	git checkout master
	git pull origin master
	git merge dev
	git push origin master
	git checkout dev

# Test target: run tests and perform code checks
test: $(OUTPUT_DIR)
	test -z "$(gofmt -l .)"
	test -z "$(goimports -local $(LOCAL_PACKAGE) -d .)"
	../.github/test.sh

# HTML target: generate HTML coverage report
html:
	go tool cover -html=c.out
