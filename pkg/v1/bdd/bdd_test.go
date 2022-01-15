package bdd_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/pkg/v1/bdd"
)

func TestBDD(t *testing.T) {
	f := bdd.NewFeature(t, "bdd")

	const expCalled = 4

	var called int
	inc := func() {
		called++
	}

	f.Scenario("simple", func(t *testing.T, f *bdd.Feature) {
		f.Given("given called", inc)
		f.When("when called", inc)
		f.And("and called", inc)
		f.Then("then called", inc)
	})

	if called != expCalled {
		t.Fatal(called)
	}
}
