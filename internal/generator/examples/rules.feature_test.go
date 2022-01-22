package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/pkg/v1/bdd"
)

func TestHighlander(t *testing.T) {
	f := bdd.NewFeature(t, "Highlander")

	f.Rule("There can be only One", func(t *testing.T, f *bdd.Feature) {

		f.Background("", func(t *testing.T, f *bdd.Feature) {
			f.Given("I have overdue tasks", func() {

			})
		})

		f.Example("Only One -- More than one alive", func(t *testing.T, f *bdd.Feature) {
			f.Given("there are 3 ninjas", func() {

			})
			f.And("there are more than one ninja alive", func() {

			})
			f.When("2 ninjas meet, they will fight", func() {

			})
			f.Then("one ninja dies (but not me)", func() {

			})
			f.And("there is one ninja less alive", func() {

			})

		})

		f.Example("Only One -- One alive", func(t *testing.T, f *bdd.Feature) {
			f.Given("there is only 1 ninja alive", func() {

			})
			f.Then("he (or she) will live forever ;-)", func() {

			})

		})

	})

	f.Rule("There can be Two (in some cases)", func(t *testing.T, f *bdd.Feature) {

		f.Example("Two -- Dead and Reborn as Phoenix", func(t *testing.T, f *bdd.Feature) {

		})

	})

}
