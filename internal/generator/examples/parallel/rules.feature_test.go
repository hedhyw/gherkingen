package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/pkg/bdd"
)

func TestHighlander(t *testing.T) {
	t.Parallel()

	f := bdd.NewFeature(t, "Highlander")

	f.Rule("There can be only One", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

		background := func(t *testing.T, f *bdd.Feature) interface{} {
			/* TODO: Feel free to modify return value(s). */
			f.Given("I have overdue tasks", func() {

			})

			return nil
		}

		f.Example("Only One -- More than one alive", func(t *testing.T, f *bdd.Feature) {
			t.Parallel()

			f.Given("there are 3 ninjas", func() {
				_ = background(t, f)

			})
			f.And("there are more than one ninja alive", func() {
				_ = background(t, f)

			})
			f.When("2 ninjas meet, they will fight", func() {
				_ = background(t, f)

			})
			f.Then("one ninja dies (but not me)", func() {
				_ = background(t, f)

			})
			f.And("there is one ninja less alive", func() {
				_ = background(t, f)

			})
		})
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

	f.Rule("There can be Two (in some cases)", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

		f.Example("Two -- Dead and Reborn as Phoenix", func(t *testing.T, f *bdd.Feature) {
			t.Parallel()

		})
	})
}
