#!/bin/sh

for f in internal/generator/examples/*.feature; do
  go run cmd/gherkingen/main.go \
    -package examples_test \
    -format go \
    $f \
    > ${f}_test.go
done
