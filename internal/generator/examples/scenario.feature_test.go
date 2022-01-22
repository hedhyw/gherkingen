package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/pkg/v1/bdd"
)

func TestSomeTerseYetDescriptiveTextOfWhatIsDesired(t *testing.T) {
	f := bdd.NewFeature(t, "Some terse yet descriptive text of what is desired")

	/*   In order to realize a named business value
	As an explicit system actor
	I want to gain some beneficial outcome which furthers the goal */

	f.Scenario("Some determinable business situation", func(t *testing.T, f *bdd.Feature) {
		f.Given("some precondition", func() {

		})
		f.And("some other precondition", func() {

		})
		f.When("some action by the actor", func() {

		})
		f.And("some other action", func() {

		})
		f.And("yet another action", func() {

		})
		f.Then("some testable outcome is achieved", func() {

		})
		f.And("something else we can check happens too", func() {

		})

	})

}
