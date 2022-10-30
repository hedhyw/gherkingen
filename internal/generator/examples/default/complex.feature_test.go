package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/pkg/bdd"
)

func TestNestedBackground(t *testing.T) {
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
		_ = background(t, f)

		f.Given("I am logged in as Dr. Bill", func() {

		})
		f.When("I try to post to \"Expensive Therapy\"", func() {

		})
		f.Then("I should see \"Your article was published.\"", func() {

		})
	})

	f.Rule("There can be only One", func(_ *testing.T, f *bdd.Feature) {
		background := func(t *testing.T, f *bdd.Feature) interface{} {
			/* TODO: Feel free to modify return value(s). */
			f.Given("I have overdue tasks", func() {

			})

			return nil
		}

		f.Example("Only One -- One alive", func(t *testing.T, f *bdd.Feature) {
			_ = background(t, f)

			f.Given("there is only 1 ninja alive", func() {

			})
			f.Then("he (or she) will live forever ;-)", func() {

			})
		})
	})
}
