package bdd

import (
	"testing"
)

type Feature struct {
	*testing.T
}

func NewFeature(t *testing.T, name string) *Feature {
	t.Helper()
	t.Log("Feature: " + name)

	return &Feature{
		T: t,
	}
}

func (f Feature) Scenario(name string, fn func(t *testing.T, f *Feature)) {
	f.T.Helper()

	f.T.Run(name, func(t *testing.T) {
		t.Helper()
		t.Log("Scenario: " + name)

		fn(t, &Feature{T: t})
	})
}

func (f Feature) Given(given string, fn func()) {
	f.T.Helper()
	f.T.Log("Given: ", given)

	if fn != nil {
		fn()
	}
}

func (f Feature) And(and string, fn func()) {
	f.T.Log("And: ", and)

	if fn != nil {
		fn()
	}
}

func (f Feature) When(when string, fn func()) {
	f.T.Log("When: ", when)

	if fn != nil {
		fn()
	}
}

func (f Feature) Then(then string, fn func()) {
	f.T.Log("Then: ", then)

	if fn != nil {
		fn()
	}
}
