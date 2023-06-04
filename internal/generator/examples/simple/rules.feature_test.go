package examples_test

import (
	"testing"
)

func TestHighlander(t *testing.T) {

	t.Run("There can be only One", func(_ *testing.T) {
		type backgroundData struct{}

		background := func(t *testing.T) backgroundData {
			// Given I have overdue tasks.

			return backgroundData{}
		}

		t.Run("Only One -- More than one alive", func(t *testing.T) {
			_ = background(t)

			// Given there are 3 ninjas.

			// And there are more than one ninja alive.

			// When 2 ninjas meet, they will fight.

			// Then one ninja dies (but not me).

			// And there is one ninja less alive.

		})
		t.Run("Only One -- One alive", func(t *testing.T) {
			_ = background(t)

			// Given there is only 1 ninja alive.

			// Then he (or she) will live forever ;-).

		})
	})

	t.Run("There can be Two (in some cases)", func(_ *testing.T) {
		t.Run("Two -- Dead and Reborn as Phoenix", func(t *testing.T) {
		})
	})
}
