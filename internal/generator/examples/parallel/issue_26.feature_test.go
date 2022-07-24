package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/pkg/bdd"
)

func TestIssueExample(t *testing.T) {
	t.Parallel()

	f := bdd.NewFeature(t, "Issue example")

	f.Scenario("Just a hello world", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

		type testCase struct {
			Name string `field:"<name>"`
		}

		testCases := map[string]testCase{
			"hello_world": {"hello world"},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			t.Parallel()

		})
	})
}
