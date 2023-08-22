# Makefile for suprelog
# Author: Pokeya
# Version: 1.0

# Set phony targets
.PHONY: merge gen help

# Define variable
LOCAL_PACKAGE := github.com/pokeyaro/gopkg/suprelog

# Default target: print available commands
.PHONY: help
.DEFAULT_GOAL := help
help:
	@echo "Usage:"
	@echo "  make gen           Generate code using code generation tool"
	@echo "  make merge         Merge dev branch into master branch"

gen:
	go generate ./....

# Merge target: merge dev branch into master branch
merge:
	git checkout master
	git pull origin master
	git merge dev
	git push origin master
	git checkout dev
