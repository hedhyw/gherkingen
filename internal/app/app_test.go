package app_test

import (
	"bytes"
	"flag"
	"testing"

	"github.com/hedhyw/gherkingen/v2/internal/app"
	"github.com/hedhyw/gherkingen/v2/pkg/bdd"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testVersion = "0.0.1"

func TestApplicationCommandLineTool(t *testing.T) {
	f := bdd.NewFeature(t, "Application command line tool")

	f.Scenario("User wants to generate the output in given format", func(_ *testing.T, f *bdd.Feature) {
		type testCase struct {
			Feature   string `field:"<feature>"`
			Format    string `field:"<format>"`
			Assertion string `field:"<assertion>"`
		}

		testCases := map[string]testCase{
			"app.feature_go_does":           {"app.feature", "go", "does"},
			"app.feature_json_does":         {"app.feature", "json", "does"},
			"app.feature_raw_does":          {"app.feature", "raw", "does"},
			"app.feature_invalid_does not":  {"app.feature", "invalid", "does not"},
			"notfound.feature_raw_does not": {"notfound.feature", "raw", "does not"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			arguments := []string{}
			f.When("<format> is given", func() {
				arguments = append(arguments, "-format", tc.Format)
			})
			f.And("<feature> is provided", func() {
				arguments = append(arguments, tc.Feature)
			})
			f.Then("the output should be generated", func() {
				runApp(t, arguments, tc.Assertion == "does")
			})
		})
	})

	f.Scenario("User wants to see usage information", func(_ *testing.T, f *bdd.Feature) {
		type testCase struct {
			Flag string `field:"<flag>"`
		}

		testCases := map[string]testCase{
			"--help": {"--help"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			arguments := []string{}
			f.When("<flag> is provided", func() {
				arguments = append(arguments, tc.Flag)
			})
			f.Then("usage should be printed", func() {
				runApp(t, arguments, true)
			})
		})
	})

	f.Scenario("User wants to list built-in templates", func(_ *testing.T, f *bdd.Feature) {
		type testCase struct {
			Flag string `field:"<flag>"`
		}

		testCases := map[string]testCase{
			"--list": {"--list"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			arguments := []string{}
			f.When("<flag> is provided", func() {
				arguments = append(arguments, tc.Flag)
			})
			f.Then("templates should be printed", func() {
				runApp(t, arguments, true)
			})
		})
	})

	f.Scenario("User wants to use custom template", func(_ *testing.T, f *bdd.Feature) {
		type testCase struct {
			Feature  string `field:"<feature>"`
			Template string `field:"<template>"`
		}

		testCases := map[string]testCase{
			"app.feature_../assets/std.struct.v1.go.tmpl": {"app.feature", "../assets/std.struct.v1.go.tmpl"},
			"app.feature_@/std.struct.v1.go.tmpl":         {"app.feature", "@/std.struct.v1.go.tmpl"},
			"app.feature_@/std.simple.v1.go.tmpl":         {"app.feature", "@/std.simple.v1.go.tmpl"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			arguments := []string{}
			f.And("<template> is provided", func() {
				arguments = append(arguments, "-template", tc.Template)
			})
			f.When("<feature> is provided", func() {
				arguments = append(arguments, tc.Feature)
			})
			f.Then("the output should be generated", func() {
				runApp(t, arguments, true)
			})
		})
	})

	f.Scenario("User wants to set custom package", func(_ *testing.T, f *bdd.Feature) {
		type testCase struct {
			Package string `field:"<package>"`
		}

		testCases := map[string]testCase{
			"app_test":     {"app_test"},
			"example_test": {"example_test"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			arguments := []string{}
			f.When("<package> is provided", func() {
				arguments = append(arguments, "-package", tc.Package, "app.feature")
			})
			f.Then("the output should contain <package>", func() {
				out := runApp(t, arguments, true)
				assert.Contains(t, out, tc.Package)
			})
		})
	})

	f.Scenario("User wants to generate a permanent json output", func(_ *testing.T, f *bdd.Feature) {
		type testCase struct {
			TheSameIDs bool `field:"<TheSameIDs>"`
		}

		testCases := map[string]testCase{
			"true":  {true},
			"false": {false},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			arguments := []string{}
			f.When("format is json", func() {
				arguments = append(arguments, "-format", "json")
			})
			f.And("-permanent-ids is <TheSameIDs>", func() {
				if tc.TheSameIDs {
					arguments = append(arguments, "-permanent-ids")
				}
			})
			f.Then("calling generation twice will produce the same output <TheSameIDs>", func() {
				arguments = append(arguments, "app.feature")

				firstOut := runApp(t, arguments, true)
				secondOut := runApp(t, arguments, true)
				if tc.TheSameIDs {
					assert.Equal(t, firstOut, secondOut)
				} else {
					assert.NotEqual(t, firstOut, secondOut)
				}
			})
		})
	})

	f.Scenario("User gives an invalid flag", func(t *testing.T, f *bdd.Feature) {
		arguments := []string{}
		f.When("flag -invalid is provided", func() {
			arguments = append(arguments, "-invalid")
		})
		f.Then("a generation failed", func() {
			arguments = append(arguments, "app.feature")
			runApp(t, arguments, false)
		})
	})

	f.Scenario("User wants to know version", func(_ *testing.T, f *bdd.Feature) {
		type testCase struct {
			Flag string `field:"<flag>"`
		}

		testCases := map[string]testCase{
			"--version": {"--version"},
			"-version":  {"-version"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			arguments := []string{}

			f.When("<flag> is provided", func() {
				arguments = append(arguments, tc.Flag)
			})
			f.Then("version is printed", func() {
				out := runApp(t, arguments, true)
				assert.Contains(t, out, testVersion)
			})
		})
	})

	f.Scenario("User specifies a file, but the file is not found", func(_ *testing.T, f *bdd.Feature) {
		type testCase struct {
			Feature  string `field:"<feature>"`
			Template string `field:"<template>"`
		}

		testCases := map[string]testCase{
			"app.feature_not_found": {"app.feature", "not_found"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			arguments := []string{}
			f.When("inexistent <template> is provided", func() {
				arguments = append(arguments, "-template", tc.Template)
			})
			f.And("<feature> is provided", func() {
				arguments = append(arguments, tc.Feature)
			})
			f.Then("the user receives an error", func() {
				runApp(t, arguments, false)
			})
		})
	})

	f.Scenario("User wants to run tests in parallel", func(t *testing.T, f *bdd.Feature) {
		arguments := []string{}
		f.When("`-go-parallel` is provided", func() {
			arguments = append(arguments, "-go-parallel")
		})
		f.And("`app.feature` is given", func() {
			arguments = append(arguments, "../generator/examples/scenario.feature")
		})
		f.Then("generated code contains `t.Parallel()`", func() {
			assert.Contains(t, runApp(t, arguments, true), "t.Parallel()")
		})
	})

	f.Scenario("User wants to run tests sequentially", func(t *testing.T, f *bdd.Feature) {
		arguments := []string{}
		f.When("`-go-parallel` is provided", func() {
			// Go on.
		})
		f.And("`app.feature` is given", func() {
			arguments = append(arguments, "../generator/examples/scenario.feature")
		})
		f.Then("generated code doesn't contain `t.Parallel()`", func() {
			assert.NotContains(t, runApp(t, arguments, true), "t.Parallel()")
		})
	})
}

func runApp(tb testing.TB, arguments []string, ok bool) string {
	tb.Helper()

	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)

	var buf bytes.Buffer
	err := app.Run(arguments, &buf, testVersion)
	if ok {
		require.NoError(tb, err)

		gotLen := buf.Len()
		assert.NotZero(tb, gotLen)
	} else {
		require.Error(tb, err)
	}

	return buf.String()
}
