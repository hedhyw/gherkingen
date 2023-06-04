package examples_test

import (
	"testing"
)

func TestNestedBackground(t *testing.T) {
	t.Parallel()

	type backgroundData struct{}

	background := func(t *testing.T) backgroundData {
		// Given a global administrator named "Greg".

		// And a blog named "Greg's anti-tax rants".

		// And a customer named "Dr. Bill".

		// And a blog named "Expensive Therapy" owned by "Dr. Bill".

		return backgroundData{}
	}

	t.Run("Dr. Bill posts to his own blog", func(t *testing.T) {
		t.Parallel()

		_ = background(t)

		// Given I am logged in as Dr. Bill.

		// When I try to post to "Expensive Therapy".

		// Then I should see "Your article was published.".

	})

	t.Run("There can be only One", func(t *testing.T) {
		t.Parallel()

		type backgroundData struct{}

		background := func(t *testing.T) backgroundData {
			// Given I have overdue tasks.

			return backgroundData{}
		}

		t.Run("Only One -- One alive", func(t *testing.T) {
			t.Parallel()

			_ = background(t)

			// Given there is only 1 ninja alive.

			// Then he (or she) will live forever ;-).

		})
	})
}
