package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/pkg/v1/bdd"
)

func TestApplicationCommandLineTool(t *testing.T) {
	f := bdd.NewFeature(t, "Application command line tool")

	f.Scenario("User wants to see usage information", func(t *testing.T, f *bdd.Feature) {
		type testCase struct {
			Flag       string `field:"<flag>"`
			ExitStatus string `field:"<exit_status>"`
			Printed    string `field:"<printed>"`
		}

		testCases := map[string]testCase{
			"--help_0_true":    {"--help", "0", "true"},
			"-help_0_true":     {"-help", "0", "true"},
			"-invalid_1_false": {"-invalid", "1", "false"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			f.When("flag <flag> is provided", func() {

			})
			f.Then("usage should be printed <printed>", func() {

			})
			f.And("exit status should be <exit_status>", func() {

			})
		})
	})

}
