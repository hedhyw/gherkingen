GOLANG_CI_LINT_VER:=v1.43.0

lint: bin/golangci-lint
	./bin/golangci-lint run
.PHONY: lint

test:
	go test -covermode=count -coverprofile=coverage.out ./...
.PHONY: test.coverage

readme.examples:
	go run cmd/gherkingen/main.go readme.feature.example > readme.go.example
.PHONY: readme

bin/golangci-lint:
	curl \
		-sSfL \
		https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s $(GOLANG_CI_LINT_VER)
