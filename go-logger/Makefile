# Makefile for go-logger
# Author: Pokeya
# Version: 1.0

# Set phony targets
.PHONY: merge help

# Default target: print available commands
.PHONY: help
.DEFAULT_GOAL := help
help:
	@echo "Usage:"
	@echo "  make merge         Merge dev branch into master branch"

# Merge target: merge dev branch into master branch
merge:
	git checkout master
	git pull origin master
	git merge dev
	git push origin master
	git checkout dev
