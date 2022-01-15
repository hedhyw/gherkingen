package app_test

import (
	"bytes"
	"flag"
	"testing"

	"github.com/hedhyw/gherkingen/internal/app"
	"github.com/hedhyw/gherkingen/pkg/v1/bdd"
)

func runApp(tb testing.TB, arguments []string, ok bool) {
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
}

func TestApplicationCommandLineTool(t *testing.T) {
	f := bdd.NewFeature(t, "Application command line tool")

	f.Scenario("User wants to generate the output in given format", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			feature   string
			format    string
			assertion string
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

			t.Run(name, func(t *testing.T) {
				t.Logf("TestCase: %+v", tc)

				arguments := []string{}
				f.When("<format> is given", func() {
					arguments = append(arguments, "-format", tc.format)
				})
				f.And("<feature> is provided", func() {
					arguments = append(arguments, tc.feature)
				})
				f.Then("the output should be generated", func() {
					runApp(t, arguments, tc.assertion == "does")
				})
			})
		}
	})

	f.Scenario("User wants to see usage information", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			flag string
		}

		testCases := map[string]testCase{
			"--help": {"--help"},
		}

		for name, tc := range testCases {
			name, tc := name, tc

			t.Run(name, func(t *testing.T) {
				t.Logf("TestCase: %+v", tc)

				arguments := []string{}
				f.When("<flag> is provided", func() {
					arguments = append(arguments, tc.flag)
				})
				f.Then("usage should be printed", func() {
					runApp(t, arguments, true)
				})
			})
		}
	})

	f.Scenario("User wants to list built-in templates", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			flag string
		}

		testCases := map[string]testCase{
			"--list": {"--list"},
		}

		for name, tc := range testCases {
			name, tc := name, tc

			t.Run(name, func(t *testing.T) {
				t.Logf("TestCase: %+v", tc)

				arguments := []string{}
				f.When("<flag> is provided", func() {
					arguments = append(arguments, tc.flag)
				})
				f.Then("templates should be printed", func() {
					runApp(t, arguments, true)
				})
			})
		}
	})

	f.Scenario("User wants to use custom template", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			feature  string
			template string
		}

		testCases := map[string]testCase{
			"app.feature_../assets/std.args.go.tmpl": {"app.feature", "../assets/std.args.go.tmpl"},
			"app.feature_@/std.args.go.tmpl":         {"app.feature", "@/std.args.go.tmpl"},
			"app.feature_@/std.struct.go.tmpl":       {"app.feature", "@/std.struct.go.tmpl"},
		}

		for name, tc := range testCases {
			name, tc := name, tc

			t.Run(name, func(t *testing.T) {
				t.Logf("TestCase: %+v", tc)

				arguments := []string{}
				f.And("<template> is provided", func() {
					arguments = append(arguments, "-template", tc.template)
				})
				f.When("<feature> is provided", func() {
					arguments = append(arguments, tc.feature)
				})
				f.Then("the output should be generated", func() {
					runApp(t, arguments, true)
				})
			})
		}
	})
}
