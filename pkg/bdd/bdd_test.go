package bdd_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v3/pkg/bdd"
)

func TestBDDTestCases(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		f := bdd.NewFeature(t, "bdd")

		type testCase struct {
			Inc int
		}

		const expSum = 20
		testCases := map[string]testCase{
			"by_7":  {7},
			"by_13": {13},
		}

		var gotSum int

		f.TestCases(testCases, func(_ *testing.T, _ *bdd.Feature, tc testCase) {
			gotSum += tc.Inc
		})

		if gotSum != expSum {
			t.Fatalf("Exp: %d, Got: %d", expSum, gotSum)
		}
	})
}

func TestBDD(t *testing.T) {
	f := bdd.NewFeature(t, "bdd")

	const expCalled = 6

	var called int
	inc := func() {
		called++
	}

	f.Rule("rule", func(_ *testing.T, f *bdd.Feature) {
		f.Background("background", func(_ *testing.T, f *bdd.Feature) {
			f.Then("then", inc)
		})

		f.Scenario("simple", func(_ *testing.T, f *bdd.Feature) {
			tc := struct {
				Fn string `field:"<field>"`
			}{
				Fn: "FUNC",
			}

			f.Example("example", func(_ *testing.T, f *bdd.Feature) {
				f.TestCase("testCase", tc, func(t *testing.T, f *bdd.Feature) {
					f.Given("given <field> called", inc)
					f.But("but <field> called", inc)
					f.When("when <field> called", inc)
					f.And("and <field> called", inc)
					f.Then("then <field> called", inc)

					expRecords := [...]string{
						"Feature: bdd",
						"\tRule: rule",
						"\t\tScenario: simple",
						"\t\t\tExample: example",
						"\t\t\t\t# TestCase: {Fn:FUNC}",
						"\t\t\t\tGiven given FUNC called",
						"\t\t\t\tBut but FUNC called",
						"\t\t\t\tWhen when FUNC called",
						"\t\t\t\tAnd and FUNC called",
						"\t\t\t\tThen then FUNC called",
					}

					records := f.LogRecords()
					if len(records) != len(expRecords) {
						t.Fatalf("Got records (%d): %+v", len(records), records)
					}

					for i, er := range expRecords {
						if records[i] != er {
							t.Fatalf("Not matched: %q and %q", records[i], er)
						}
					}
				})
			})
		})
	})

	if called != expCalled {
		t.Fatal(called)
	}
}
