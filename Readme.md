# gherkingen

![Version](https://img.shields.io/github/v/tag/hedhyw/gherkingen)
[![Build Status](https://app.travis-ci.com/hedhyw/gherkingen.svg?branch=main)](https://app.travis-ci.com/github/hedhyw/gherkingen)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedhyw/gherkingen)](https://goreportcard.com/report/github.com/hedhyw/gherkingen)
[![Coverage Status](https://coveralls.io/repos/github/hedhyw/gherkingen/badge.svg?branch=master)](https://coveralls.io/github/hedhyw/gherkingen?branch=master)

**It's a Behaviour Driven Development (BDD) tests generator for Golang.**

<img alt="Golang logo" src="https://raw.githubusercontent.com/rfyiamcool/golang_logo/master/png/golang_55.png" height="200" />



It accepts `*.feature` [Cucumber/Gherkin](https://cucumber.io/docs/gherkin/reference/) files and generates a test boilerplate. All that remains is to change the tests a little.

The generator is very customizable, it is possible to customize an output for any golang [testing framework](#frameworks-support) or even for [another language](#language-support).

# What is for?
## Simple example
**Given** [feature](readme.feature.example) [[reference](https://cucumber.io/docs/gherkin/reference/)]:
```feature
Feature: Application command line tool
  Scenario: User wants to see usage information
    When <flag> is provided
    Then usage should be printed
    Examples:
    | <flag> |
    | --help |
    | -help  |
```

**Then** this generator writes a [golang](readme.go.example) output:

```go
func TestApplicationCommandLineTool(t *testing.T) {
	f := bdd.NewFeature(t, "Application command line tool")

	f.Scenario("User wants to see usage information", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			flag string
		}

		testCases := map[string]testCase{
			"--help": {"--help"},
			"-help":  {"-help"},
		}

		for name, tc := range testCases {
			name, tc := name, tc

			t.Run(name, func(t *testing.T) {
				t.Logf("TestCase: %+v", tc)
				f.When("<flag> is provided", func() {

				})
				f.Then("usage should be printed", func() {

				})
			})
		}
	})
}
```
## More advanced example

See [internal/app/app.feature](internal/app/app.feature) and [internal/app/app_test.go](internal/app/app_test.go).

# Install

Run:
```
go install github.com/hedhyw/gherkingen/cmd/gherkingen@latest
```

# Usage
## Simple usage

For generating test output, simply run:

```
gherkingen EXAMPLE.feature
```

## More advanced usage

### Generating test output with custom options
```
gherkingen \
    -format go \
    -template my_template.tmpl \
    EXAMPLE.feature
```
### Listing internal templates
```
gherkingen -list
```

### Help
```
gherkingen --help

Usage of gherkingen [FEATURE_FILE]:
  -format string
        output format: json, go, raw (default "go")
  -help
        print usage
  -list
        list internal templates
  -template string
        template file (default "@/std.struct.go.tmpl")
```

# Output customization

## Custom templates
You can provide your own template, it can be based on [internal/assets/std.args.go.tmpl](internal/assets/std.args.go.tmpl). In the command-line tool specify the template
using `-template` flag: `gherkingen -template example.tmpl raw example.feature`

## Frameworks support
It is possible to integrate with any BDD-testing fraemwork. Feel free to
create a pull request for supporting templates for them. For this:
1. Create a template `internal/assets/SOME_NAME.go.tmpl`.
2. Add it to the test `TestOpenTemplate` in the file [internal/assets/assets_test.go](internal/assets/assets_test.go).
3. Check: `make lint test`.

## Language support

Templates are very customizable, so you can even generate non-golang code. In the command-line tool specify `raw` format using `-format` flag:
`gherkingen -format raw example.feature`.

## License

See [License](License).
