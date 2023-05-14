package examples_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/pkg/bdd"
)

func TestExampleIssue27Single(t *testing.T) {
	f := bdd.NewFeature(t, "Example Issue 27 Single")

	/* Hello world. */

	f.Example("Single comment", func(t *testing.T, f *bdd.Feature) {
	})
}
