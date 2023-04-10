package examples_test

import (
	"testing"
)

func TestIssueExample(t *testing.T) {
	t.Parallel()

	t.Run("Just a hello world", func(t *testing.T) {
		t.Parallel()

		type testCase struct {
			Name string `field:"<name>"`
		}

		testCases := map[string]testCase{
			"hello_world": {"hello world"},
		}

		for name, tc := range testCases {
			tc := tc
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				_ = tc // TODO: Use and remove.
			})
		}
	})
}
