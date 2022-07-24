#!/bin/sh

for f in internal/generator/examples/*.feature; do
  dir=$(dirname $f)
  base=$(basename $f)
  go run cmd/gherkingen/main.go \
    -package examples_test \
    -permanent-ids \
    -format go \
    $f \
    > ${dir}/default/${base}_test.go
  go run cmd/gherkingen/main.go \
    -package examples_test \
    -permanent-ids \
    -format go \
    -go-parallel \
    $f \
    > ${dir}/parallel/${base}_test.go
  go run cmd/gherkingen/main.go \
    -package examples_test \
    -permanent-ids \
    -format json \
    $f \
    > ${dir}/default/${base}.json
  go run cmd/gherkingen/main.go \
    -package examples_test \
    -permanent-ids \
    -format json \
    -go-parallel \
    $f \
    > ${dir}/parallel/${base}.json
done
