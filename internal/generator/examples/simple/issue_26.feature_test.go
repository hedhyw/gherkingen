package examples_test

import (
	"testing"
)

func TestIssueExample(t *testing.T) {

	t.Run("Just a hello world", func(_ *testing.T) {
		type testCase struct {
			Name string `field:"<name>"`
		}

		testCases := map[string]testCase{
			"hello_world": {"hello world"},
		}

		for name, testCase := range testCases {
			t.Run(name, func(t *testing.T) {
				_ = testCase // TODO: Use and remove.
			})
		}
	})
}
