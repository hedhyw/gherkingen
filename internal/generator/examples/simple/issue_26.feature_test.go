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

		for name, testCase := range testCases {
			testCase := testCase

			t.Run(name, func(t *testing.T) {
				t.Parallel()

				_ = testCase // TODO: Use and remove.
			})
		}
	})
}
