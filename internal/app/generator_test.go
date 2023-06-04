package app

import (
	"testing"

	"github.com/hedhyw/gherkingen/v3/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestDetectFormat(t *testing.T) {
	t.Parallel()

	testCases := [...]struct {
		Filename string
		Expected model.Format
	}{{
		Filename: "example.go",
		Expected: model.FormatGo,
	}, {
		Filename: "example.rb",
		Expected: model.FormatRaw,
	}, {
		Filename: "example.json",
		Expected: model.FormatRaw,
	}, {
		Filename: "example.GO",
		Expected: model.FormatGo,
	}, {
		Filename: "example.go.tmpl",
		Expected: model.FormatGo,
	}, {
		Filename: "example.rb.tmpl",
		Expected: model.FormatRaw,
	}}

	for _, tc := range testCases {
		actual := detectFormat(tc.Filename)
		assert.Equal(t, tc.Expected, actual, tc.Filename)
	}
}
