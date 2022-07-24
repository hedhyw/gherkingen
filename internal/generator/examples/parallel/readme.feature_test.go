package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/pkg/bdd"
)

func TestApplicationCommandLineTool(t *testing.T) {
	t.Parallel()

	f := bdd.NewFeature(t, "Application command line tool")

	f.Scenario("User wants to see usage information", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

		type testCase struct {
			Flag       string `field:"<flag>"`
			ExitStatus int    `field:"<exit_status>"`
			Printed    bool   `field:"<printed>"`
		}

		testCases := map[string]testCase{
			"--help_0_true":    {"--help", 0, true},
			"-help_0_true":     {"-help", 0, true},
			"-invalid_1_false": {"-invalid", 1, false},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			t.Parallel()

			f.When("flag <flag> is provided", func() {

			})
			f.Then("usage should be printed <printed>", func() {

			})
			f.And("exit status should be <exit_status>", func() {

			})
		})
	})
}
