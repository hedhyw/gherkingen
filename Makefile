GOLANG_CI_LINT_VER:=v1.45.0
OUT_BIN?=${PWD}/bin/gherkingen

build:
	go build -o ${OUT_BIN} cmd/gherkingen/main.go
.PHONY: build

lint: bin/golangci-lint
	./bin/golangci-lint run
.PHONY: lint

test:
	go test -coverpkg=./... -covermode=count -coverprofile=coverage.out ./...
.PHONY: test

generate:
	sh scripts/examples.sh
.PHONY: generate

check.generate: generate
	git diff --exit-code internal/generator/examples
.PHONY: check.generate

vendor:
	go mod tidy
	go mod vendor
.PHONY: vendor

bin/golangci-lint:
	curl \
		-sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s $(GOLANG_CI_LINT_VER)
