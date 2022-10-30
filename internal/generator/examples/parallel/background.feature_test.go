package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/pkg/bdd"
)

func TestMultipleSiteSupport(t *testing.T) {
	t.Parallel()

	f := bdd.NewFeature(t, "Multiple site support")

	/*
		Only blog owners can post to a blog, except administrators,
		who can post to all blogs.
	*/

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

		_ = background(t, f)

		f.Given("I am logged in as Dr. Bill", func() {

		})
		f.When("I try to post to \"Expensive Therapy\"", func() {

		})
		f.Then("I should see \"Your article was published.\"", func() {

		})
	})

	f.Scenario("Dr. Bill tries to post to somebody else's blog, and fails", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

		_ = background(t, f)

		f.Given("I am logged in as Dr. Bill", func() {

		})
		f.When("I try to post to \"Greg's anti-tax rants\"", func() {

		})
		f.Then("I should see \"Hey! That's not your blog!\"", func() {

		})
	})

	f.Scenario("Greg posts to a client's blog", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

		_ = background(t, f)

		f.Given("I am logged in as Greg", func() {

		})
		f.When("I try to post to \"Expensive Therapy\"", func() {

		})
		f.Then("I should see \"Your article was published.\"", func() {

		})
	})
}
