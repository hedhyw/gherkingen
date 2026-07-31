# AGENTS.md

Guidance for AI coding agents working in this repository.

## What this project is

`gherkingen` is a command-line generator of Behaviour Driven Development (BDD)
test boilerplate for Go. It parses a Cucumber/Gherkin `*.feature` file and
renders a Go test skeleton (table-driven, parallel by default) through a
`text/template`. Output is customizable: any Go testing framework, or even
non-Go languages, can be targeted by supplying a custom template with
`-format raw`. Module path: `github.com/hedhyw/gherkingen/v4` (binary in
`cmd/gherkingen`).

## CLI usage

```sh
gherkingen [flags] FEATURE_FILE
```

Key flags (see `internal/app/app.go` for the authoritative list):

- `-format` — output format: `autodetect` (default), `json`, `go`, `raw`.
- `-template` — template file; `@/` prefix refers to embedded templates
  (default `@/std.simple.v1.go.tmpl`).
- `-package` — generated package name (default `generated_test`).
- `-language` — natural language of the feature (default `en`); may also be
  inferred from a `<name>.<lang>.feature` file name.
- `-permanent-ids` — deterministic UUIDs, same input gives same output.
- `-disable-go-parallel` — do not emit `t.Parallel()`.
- `-list` / `-languages` / `-version` / `-help` — informational.

There is no public Go API: all packages live under `internal/`.

## Repository layout

- `cmd/gherkingen/main.go` — entry point, calls `internal/app.Run`.
- `internal/app` — CLI wiring: flag parsing, template resolution
  (`templates.go`), language handling, `app.feature`/`app_test.go` BDD tests.
- `internal/generator` — renders the parsed document: `golang.go` (Go
  templates), `json.go`, `raw.go`.
- `internal/generator/examples` — example `*.feature` inputs and their
  GENERATED outputs in `simple/` and `simpleparallel/` (see below).
- `internal/model` — data model passed to templates (`TemplateData`, Gherkin
  document types, format enums).
- `internal/assets` — embedded templates (`//go:embed *.tmpl`), currently
  `std.simple.v1.go.tmpl`.
- `internal/docplugin` — plugins that inject extra data into the document
  before rendering; `goplugin` adds Go types/aliases, `multiplugin` combines
  plugins.
- `scripts/examples.sh` — regenerates the example outputs.
- `assets/` — README images only.

## Build, test, lint

```sh
make build           # builds ./bin/gherkingen with version from git describe
make test            # go test with coverage profile (coverage.out)
make lint            # golangci-lint (version pinned via GOLANG_CI_LINT_VER in Makefile, config .golangci.json)
make generate        # regenerate internal/generator/examples outputs
make check.generate  # fails if generated examples are stale
```

CI (`.github/workflows/check.yml`) runs build, lint, `check.generate`, and
test on every PR and push to `main`.

## Generated files — do not edit by hand

`gherkingen` is itself a code generator and dogfoods its output:

- `internal/generator/examples/simple/*_test.go` and
  `internal/generator/examples/simpleparallel/*_test.go` are generated from
  the sibling `*.feature` files by `scripts/examples.sh`. After changing
  templates, the generator, or any example feature, run `make generate` and
  commit the result; CI enforces this via `make check.generate`.
- Everything else (including `internal/assets/*.go.tmpl`) is hand-written
  source.

### Adding a new template

1. Create `internal/assets/<name>.go.tmpl` (Go `text/template`; the root
   object is `internal/model.TemplateData`).
2. Register it in `TestOpenTemplate` in `internal/assets/assets_test.go`.
3. Run `make lint check.generate test`.

## Conventions

- PR titles must follow Conventional Commits (enforced by the
  `check-pr-semantic` workflow).
- Errors use `github.com/hedhyw/semerr`; tests use `stretchr/testify`.
- Dependencies are managed with Go modules only (no `vendor/` directory);
  run `go mod tidy` after changing `go.mod`.
- Releases are tag-driven (`.goreleaser.yml`, Docker image `hedhyw/gherkingen`).
