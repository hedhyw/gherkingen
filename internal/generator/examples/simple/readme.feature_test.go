package examples_test

import (
	"testing"
)

func TestApplicationCommandLineTool(t *testing.T) {

	t.Run("User wants to see usage information", func(_ *testing.T) {
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

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				_ = testCase // TODO: Use and remove.
				// When the application is started with <flag>.

				// Then usage should be printed <printed>.

				// And exit status should be <exit_status>.

			})
		}
	})
}
