package assets_test

import (
	"io/ioutil"
	"testing"

	"github.com/hedhyw/gherkingen/internal/assets"

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
		"std.struct.v1.go.tmpl",
	}

	for _, f := range files {
		f := f

		t.Run(f, func(t *testing.T) {
			t.Parallel()

			f, err := assets.OpenTemplate(f)
			require.NoError(t, err)

			defer func() { assert.NoError(t, f.Close()) }()

			data, err := ioutil.ReadAll(f)
			if assert.NoError(t, err) {
				assert.NotEmpty(t, data)
			}
		})
	}
}
