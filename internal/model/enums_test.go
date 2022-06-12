package model_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestFormats(t *testing.T) {
	t.Parallel()

	actualFormats := model.Formats()

	expFormats := [...]string{
		string(model.FormatJSON),
		string(model.FormatGo),
		string(model.FormatRaw),
	}

	assert.EqualValues(t, expFormats[:], actualFormats)
}
