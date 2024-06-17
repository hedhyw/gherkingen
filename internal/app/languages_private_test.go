package app

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunListFeatureLanguages(t *testing.T) {
	t.Parallel()

	t.Run("it produces a stable output", func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer

		require.NoError(t, runListFeatureLanguages(&buf))

		out1 := buf.String()

		buf.Reset()

		require.NoError(t, runListFeatureLanguages(&buf))

		out2 := buf.String()

		assert.Equal(t, out1, out2)
	})
}
