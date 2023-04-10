package examples_test

import (
	"testing"
)

func TestApplicationCommandLineTool(t *testing.T) {
	t.Parallel()

	t.Run("User wants to see usage information", func(t *testing.T) {
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

		for name, tc := range testCases {
			tc := tc
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				_ = tc // TODO: Use and remove.
				// When flag <flag> is provided

				// Then usage should be printed <printed>

				// And exit status should be <exit_status>

			})
		}
	})
}
