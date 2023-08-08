#!/bin/bash

set -e

module_name=$(cat go.mod | grep module | cut -d ' ' -f 2-2)
echo "module_name is $module_name"

GO111MODULE=on echo 'mode: atomic' > c.out && \
  go list ./... | xargs -n1 -I{} sh -c 'LOCAL_TEST=true go test -v -covermode=atomic -coverprofile=coverage.tmp -coverpkg=./... -parallel 1 -p 1 -count=1 -gcflags=-l {} && tail -n +2 coverage.tmp >> c.out' && \
  ( rm coverage.tmp || echo "coverage.tmp not exist")

# go tool cover -func=c.out -o coverage.txt
