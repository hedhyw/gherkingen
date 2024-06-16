#!/bin/sh

for f in internal/generator/examples/*.feature; do
  dir=$(dirname "$f")
  base=$(basename "$f")
  name=$(basename "$f" .feature)
  language="en" # The default language.

  # Check if the file name contains a hint about a specific language.
  case "$name" in
    *.*)
        language=$(echo "$name" | awk -F'.' '{print $NF}')
    ;;
  esac

  go run cmd/gherkingen/main.go \
    -package examples_test \
    -permanent-ids \
    -format go \
    -go-parallel \
    -template "@/std.simple.v1.go.tmpl" \
    -language "$language" \
    "$f" \
    > "${dir}"/simpleparallel/"${base}"_test.go
  go run cmd/gherkingen/main.go \
    -package examples_test \
    -permanent-ids \
    -format go \
    -template "@/std.simple.v1.go.tmpl" \
    -language "$language" \
    "$f" \
    > "${dir}"/simple/"${base}"_test.go
done
