package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/pkg/bdd"
)

func TestExampleIssue27Multi(t *testing.T) {
	t.Parallel()

	f := bdd.NewFeature(t, "Example Issue 27 Multi")

	/*
		Details:
		  - example 1
		  - example 2

		  - example 3
		    - example 3.1
		    - example 3.2
	*/

	f.Example("Multi-line comment with indents", func(t *testing.T, f *bdd.Feature) {
		t.Parallel()

	})
}
