package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/pkg/bdd"
)

func TestTypeDeterminatiopn(t *testing.T) {
	f := bdd.NewFeature(t, "Type determinatiopn")

	f.Scenario("All type are determinated", func(_ *testing.T, f *bdd.Feature) {
		type testCase struct {
			Bool    bool    `field:"<bool>"`
			Int     int     `field:"<int>"`
			String  string  `field:"<string>"`
			Flag    bool    `field:"<flag>"`
			Float64 float64 `field:"<float64>"`
		}

		testCases := map[string]testCase{
			"true_1_hello_-_1.0":  {true, 1, "hello", false, 1.0},
			"false_2_world_+_0.0": {false, 2, "world", true, 0.0},
		}

		f.TestCases(testCases, func(t *testing.T, f *bdd.Feature, tc testCase) {
			f.When("generator completed", func() {

			})
			f.Then("correct types are shown", func() {

			})
		})
	})
}
