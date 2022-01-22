GOLANG_CI_LINT_VER:=v1.43.0

lint: bin/golangci-lint
	./bin/golangci-lint run
.PHONY: lint

test:
	go test -covermode=count -coverprofile=coverage.out ./...
.PHONY: test.coverage

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
