package app_test

import (
	"bytes"
	"testing"

	gherkin "github.com/cucumber/gherkin/go/v28"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hedhyw/gherkingen/v4/internal/app"
)

const testVersion = "0.0.1"

func TestApplicationCommandLineTool(t *testing.T) {
	t.Parallel()

	t.Run("User wants to generate the output in given format", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Feature   string `field:"<feature>"`
			Format    string `field:"<format>"`
			Assertion string `field:"<assertion>"`
		}

		testCases := map[string]testCase{
			"app.feature_go_does":           {"app.feature", "go", "does"},
			"app.feature_json_does":         {"app.feature", "json", "does"},
			"app.feature_raw_does":          {"app.feature", "raw", "does"},
			"app.feature_invalid_does_not":  {"app.feature", "invalid", "does not"},
			"notfound.feature_raw_does_not": {"notfound.feature", "raw", "does not"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When <format> is given.
				arguments := []string{"-format", testCase.Format}

				// And <feature> is provided.
				arguments = append(arguments, testCase.Feature)

				// Then the output should be generated.
				runApp(t, arguments, testCase.Assertion == "does")
			})
		}
	})

	t.Run("User wants to see usage information", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Flag string `field:"<flag>"`
		}

		testCases := map[string]testCase{
			"--help": {"--help"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When <flag> is provided.
				arguments := []string{testCase.Flag}

				// Then usage should be printed.
				runApp(t, arguments, true)
			})
		}
	})

	t.Run("User wants to list built-in templates", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Flag string `field:"<flag>"`
		}

		testCases := map[string]testCase{
			"--list": {"--list"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When <flag> is provided.
				arguments := []string{testCase.Flag}

				// Then templates should be printed.
				runApp(t, arguments, true)
			})
		}
	})

	t.Run("User asks for a version", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Flag string `field:"<flag>"`
		}

		testCases := map[string]testCase{
			"--version": {"--version"},
			"-version":  {"-version"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When <flag> is provided.
				arguments := []string{testCase.Flag}

				// Then version is printed.
				out := runApp(t, arguments, true)
				assert.Contains(t, out, testVersion)
			})
		}
	})

	t.Run("User wants to generate the output for a feature written in a specific natural language", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Language  string `field:"<language>"`
			Feature   string `field:"<feature>"`
			Assertion string `field:"assertion"`
		}

		testCases := map[string]testCase{
			"en_../generator/examples/simple.feature_does":                  {"en", "../generator/examples/simple.feature", "does"},
			"en-pirate_../generator/examples/simple.en-pirate.feature_does": {"en-pirate", "../generator/examples/simple.en-pirate.feature", "does"},
			"unsupported_app.feature_does_not":                              {"unsupported", "app.feature", "does not"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When the <language> is given.
				arguments := []string{
					"-language",
					testCase.Language,
				}

				// And the <feature> is provided.
				arguments = append(arguments, testCase.Feature)

				// Then the output should be generated.
				runApp(t, arguments, testCase.Assertion == "does")
			})
		}
	})

	t.Run("User wants to see all supported natural languages", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Flag string `field:"<flag>"`
		}

		testCases := map[string]testCase{
			"-languages":  {"-languages"},
			"--languages": {"--languages"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When the <flag> is provided.
				arguments := []string{testCase.Flag}

				// Then the list of supported natural languages should be printed.
				out := runApp(t, arguments, true)

				dialectProvider := gherkin.DialectsBuiltin()
				dialect := dialectProvider.GetDialect(gherkin.DefaultDialect)

				assert.Contains(t, out, dialect.Name)
				assert.Contains(t, out, dialect.Language)
				assert.Contains(t, out, dialect.Native)
			})
		}
	})
}

func TestApplicationCommandLineToolCustom(t *testing.T) {
	t.Parallel()

	t.Run("User wants to use custom template", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Feature  string `field:"<feature>"`
			Template string `field:"<template>"`
		}

		testCases := map[string]testCase{
			"app.feature_../assets/std.simple.v1.go.tmpl": {"app.feature", "../assets/std.simple.v1.go.tmpl"},
			"app.feature_@/std.simple.v1.go.tmpl":         {"app.feature", "@/std.simple.v1.go.tmpl"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When <template> is provided.
				arguments := []string{"-template", testCase.Template}

				// And <feature> is provided.
				arguments = append(arguments, testCase.Feature)

				// Then the output should be generated.
				runApp(t, arguments, true)
			})
		}
	})

	t.Run("User wants to set custom package", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Package string `field:"<package>"`
		}

		testCases := map[string]testCase{
			"app_test":     {"app_test"},
			"example_test": {"example_test"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When <package> is provided.
				arguments := []string{"-package", testCase.Package, "app.feature"}

				// Then the output should contain <package>.
				out := runApp(t, arguments, true)
				assert.Contains(t, out, testCase.Package)
			})
		}
	})

	t.Run("User wants to generate a permanent json output", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			TheSameIDs bool `field:"<TheSameIDs>"`
		}

		testCases := map[string]testCase{
			"true":  {true},
			"false": {false},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When -format is json.
				arguments := []string{"-format", "json"}

				// And -permanent-ids is <TheSameIDs>.
				if testCase.TheSameIDs {
					arguments = append(arguments, "-permanent-ids")
				}

				// Then calling generation twice will produce the same output <TheSameIDs>.
				arguments = append(arguments, "app.feature")

				firstOut := runApp(t, arguments, true)
				secondOut := runApp(t, arguments, true)

				if testCase.TheSameIDs {
					assert.Equal(t, firstOut, secondOut)
				} else {
					assert.NotEqual(t, firstOut, secondOut)
				}
			})
		}
	})
}

func TestApplicationCommandLineToolFailures(t *testing.T) {
	t.Parallel()

	t.Run("User provides an invalid flag", func(t *testing.T) {
		t.Parallel()

		// When flag -invalid is provided.
		arguments := []string{
			"-invalid",
			"app.feature",
		}

		// Then an error is returned.
		runApp(t, arguments, false)
	})

	t.Run("User specifies a file, but the file is not found", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Feature  string `field:"<feature>"`
			Template string `field:"<template>"`
		}

		testCases := map[string]testCase{
			"app.feature_not_found": {"app.feature", "not_found"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				// When inexistent <template> is provided.
				arguments := []string{"-template", testCase.Template}

				// And <feature> is provided.
				arguments = append(arguments, testCase.Feature)

				// Then the user receives an error.
				runApp(t, arguments, false)
			})
		}
	})
}

func TestApplicationCommandLineToolParallel(t *testing.T) {
	t.Parallel()

	t.Run("User wants to run tests in parallel", func(t *testing.T) {
		t.Parallel()

		// When `scenario.feature` is given.
		arguments := []string{"../generator/examples/scenario.feature"}

		// Then generated code contains `t.Parallel()`.
		assert.Contains(t, runApp(t, arguments, true), "t.Parallel()")
	})

	t.Run("User wants to run tests sequentially", func(t *testing.T) {
		t.Parallel()

		// When `-disable-go-parallel` is provided.
		arguments := []string{"-disable-go-parallel"}

		// And `scenario.feature` is given.
		arguments = append(arguments, "../generator/examples/scenario.feature")

		// Then generated code doesn't contain `t.Parallel()`.
		assert.NotContains(t, runApp(t, arguments, true), "t.Parallel()")
	})
}

func runApp(tb testing.TB, arguments []string, ok bool) string {
	tb.Helper()

	tb.Log("running application with arguments", arguments)

	var buf bytes.Buffer
	err := app.Run(arguments, &buf, testVersion)
	if ok {
		require.NoError(tb, err)

		gotLen := buf.Len()
		assert.NotZero(tb, gotLen)
	} else {
		require.Error(tb, err, buf.String())
	}

	return buf.String()
}
