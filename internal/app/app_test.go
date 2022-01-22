package app_test

import (
	"bytes"
	"flag"
	"strings"
	"testing"

	"github.com/hedhyw/gherkingen/internal/app"
	"github.com/hedhyw/gherkingen/pkg/v1/bdd"
)

func runApp(tb testing.TB, arguments []string, ok bool) string {
	tb.Helper()

	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError)

	var buf bytes.Buffer
	err := app.Run(arguments, &buf)

	if ok == (err != nil) {
		tb.Errorf("Assertion failed, ok: %t, err: %s", ok, err)
	}

	if ok {
		gotLen := buf.Len()
		if gotLen == 0 {
			tb.Error("Empty output")
		}
	}

	return buf.String()
}

func TestApplicationCommandLineTool(t *testing.T) {
	f := bdd.NewFeature(t, "Application command line tool")

	f.Scenario("User wants to generate the output in given format", func(t *testing.T, f *bdd.Feature) {
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

		for name, tc := range testCases {
			name, tc := name, tc

			f.TestCase(name, tc, func(t *testing.T, f *bdd.Feature) {
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
		}
	})

	f.Scenario("User wants to see usage information", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			Flag string `field:"<flag>"`
		}

		testCases := map[string]testCase{
			"--help": {"--help"},
		}

		for name, tc := range testCases {
			name, tc := name, tc

			f.TestCase(name, tc, func(t *testing.T, f *bdd.Feature) {
				arguments := []string{}
				f.When("<flag> is provided", func() {
					arguments = append(arguments, tc.Flag)
				})
				f.Then("usage should be printed", func() {
					runApp(t, arguments, true)
				})
			})
		}
	})

	f.Scenario("User wants to list built-in templates", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			Flag string `field:"<flag>"`
		}

		testCases := map[string]testCase{
			"--list": {"--list"},
		}

		for name, tc := range testCases {
			name, tc := name, tc

			f.TestCase(name, tc, func(t *testing.T, f *bdd.Feature) {
				arguments := []string{}
				f.When("<flag> is provided", func() {
					arguments = append(arguments, tc.Flag)
				})
				f.Then("templates should be printed", func() {
					runApp(t, arguments, true)
				})
			})
		}
	})

	f.Scenario("User wants to use custom template", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			Feature  string `field:"<feature>"`
			Template string `field:"<template>"`
		}

		testCases := map[string]testCase{
			"app.feature_../assets/std.struct.v1.go.tmpl": {"app.feature", "../assets/std.struct.v1.go.tmpl"},
			"app.feature_@/std.struct.v1.go.tmpl":         {"app.feature", "@/std.struct.v1.go.tmpl"},
		}

		for name, tc := range testCases {
			name, tc := name, tc

			f.TestCase(name, tc, func(t *testing.T, f *bdd.Feature) {
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
		}
	})

	f.Scenario("User wants to set custom package", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			Package string `field:"<package>"`
		}

		testCases := map[string]testCase{
			"app_test":     {"app_test"},
			"example_test": {"example_test"},
		}

		for name, tc := range testCases {
			name, tc := name, tc

			f.TestCase(name, tc, func(t *testing.T, f *bdd.Feature) {
				arguments := []string{}
				f.When("<package> is provided", func() {
					arguments = append(arguments, "-package", name, "app.feature")
				})
				f.Then("the output should contain <package>", func() {
					out := runApp(t, arguments, true)
					if !strings.Contains(out, tc.Package) {
						t.Fatal(out)
					}
				})
			})
		}
	})
}
