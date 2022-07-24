package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/pkg/bdd"
)

func TestNestedBackground(t *testing.T) {
	t.Parallel()

	f := bdd.NewFeature(t, "Nested background")

	background := func(t *testing.T, f *bdd.Feature) interface{} {
		/* TODO: Feel free to modify return value(s). */
		f.Given("a global administrator named \"Greg\"", func() {

		})
		f.And("a blog named \"Greg's anti-tax rants\"", func() {

		})
		f.And("a customer named \"Dr. Bill\"", func() {

		})
		f.And("a blog named \"Expensive Therapy\" owned by \"Dr. Bill\"", func() {

		})

		return nil
	}

	f.Scenario("Dr. Bill posts to his own blog", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

		f.Given("I am logged in as Dr. Bill", func() {
			_ = background(t, f)

		})
		f.When("I try to post to \"Expensive Therapy\"", func() {
			_ = background(t, f)

		})
		f.Then("I should see \"Your article was published.\"", func() {
			_ = background(t, f)

		})
	})

	f.Rule("There can be only One", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

		background := func(t *testing.T, f *bdd.Feature) interface{} {
			/* TODO: Feel free to modify return value(s). */
			f.Given("I have overdue tasks", func() {

			})

			return nil
		}

		f.Example("Only One -- One alive", func(t *testing.T, f *bdd.Feature) {
			t.Parallel()

			f.Given("there is only 1 ninja alive", func() {
				_ = background(t, f)

			})
			f.Then("he (or she) will live forever ;-)", func() {
				_ = background(t, f)

			})
		})
	})
}
