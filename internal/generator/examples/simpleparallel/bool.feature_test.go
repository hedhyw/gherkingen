package examples_test

import (
	"testing"
)

func TestTypeDeterminatiopn(t *testing.T) {
	t.Parallel()

	t.Run("All type are determinated", func(t *testing.T) {
		t.Parallel()

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

		for name, testCase := range testCases {
			testCase := testCase
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				_ = testCase // TODO: Use and remove.
				// When generator completed.

				// Then correct types are shown.

			})
		}
	})
}
