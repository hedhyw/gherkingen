package assets_test

import (
	"io"
	"testing"

	"github.com/hedhyw/gherkingen/v4/internal/assets"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTemplates(t *testing.T) {
	t.Parallel()

	names, err := assets.Templates()
	if assert.NoError(t, err) {
		assert.NotEmpty(t, names)
	}
}

func TestOpenTemplate(t *testing.T) {
	t.Parallel()

	files := [...]string{
		"std.simple.v1.go.tmpl",
	}

	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			t.Parallel()

			f, err := assets.OpenTemplate(file)
			require.NoError(t, err)

			defer func() { assert.NoError(t, f.Close()) }()

			data, err := io.ReadAll(f)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, data)
			}
		})
	}
}
