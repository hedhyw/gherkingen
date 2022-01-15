package bdd

import (
	"testing"
)

// Feature of the test file.
type Feature struct {
	*testing.T
}

// NewFeature creates a new feature.
func NewFeature(t *testing.T, name string) *Feature {
	t.Helper()
	t.Log("Feature: " + name)

	return &Feature{
		T: t,
	}
}

// Scenario defines a scenario block.
func (f Feature) Scenario(name string, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.T.Run(name, func(t *testing.T) {
		t.Helper()
		t.Log("Scenario: " + name)

		fn(t, &Feature{T: t})
	})
}

// Given defines a given block.
func (f Feature) Given(given string, fn func()) {
	f.T.Helper()
	f.T.Log("Given: ", given)

	if fn != nil {
		fn()
	}
}

// And defines an and block.
func (f Feature) And(and string, fn func()) {
	f.T.Log("And: ", and)

	if fn != nil {
		fn()
	}
}

// When defines a when block.
func (f Feature) When(when string, fn func()) {
	f.T.Log("When: ", when)

	if fn != nil {
		fn()
	}
}

// Then defines a then block.
func (f Feature) Then(then string, fn func()) {
	f.T.Log("Then: ", then)

	if fn != nil {
		fn()
	}
}
