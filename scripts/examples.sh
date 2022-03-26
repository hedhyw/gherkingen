#!/bin/sh

for f in internal/generator/examples/*.feature; do
  go run cmd/gherkingen/main.go \
    -package examples_test \
    -permanent-ids \
    -format go \
    $f \
    > ${f}_test.go
  go run cmd/gherkingen/main.go \
    -package examples_test \
    -permanent-ids \
    -format json \
    $f \
    > ${f}.json
done
