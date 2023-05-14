package examples_test

import (
	"testing"
)

func TestMultipleSiteSupport(t *testing.T) {
	t.Parallel()

	/*
		Only blog owners can post to a blog, except administrators,
		who can post to all blogs.
	*/

	background := func(t *testing.T) interface{} {
		// Given a global administrator named "Greg".

		// And a blog named "Greg's anti-tax rants".

		// And a customer named "Dr. Bill".

		// And a blog named "Expensive Therapy" owned by "Dr. Bill".

		return nil // TODO: Feel free to modify return value(s).
	}

	t.Run("Dr. Bill posts to his own blog", func(t *testing.T) {
		t.Parallel()

		_ = background(t)

		// Given I am logged in as Dr. Bill.

		// When I try to post to "Expensive Therapy".

		// Then I should see "Your article was published.".

	})

	t.Run("Dr. Bill tries to post to somebody else's blog, and fails", func(t *testing.T) {
		t.Parallel()

		_ = background(t)

		// Given I am logged in as Dr. Bill.

		// When I try to post to "Greg's anti-tax rants".

		// Then I should see "Hey! That's not your blog!".

	})

	t.Run("Greg posts to a client's blog", func(t *testing.T) {
		t.Parallel()

		_ = background(t)

		// Given I am logged in as Greg.

		// When I try to post to "Expensive Therapy".

		// Then I should see "Your article was published.".

	})
}
