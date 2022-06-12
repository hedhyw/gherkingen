package multiplugin_test

import (
	"context"
	"testing"

	"github.com/hedhyw/gherkingen/internal/docplugin/goplugin"
	"github.com/hedhyw/gherkingen/internal/docplugin/multiplugin"
	"github.com/hedhyw/gherkingen/internal/model"

	"github.com/hedhyw/semerr/pkg/v1/semerr"
	"github.com/stretchr/testify/assert"
)

func TestMultiPlugin(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("name", func(t *testing.T) {
		t.Parallel()

		mp := multiplugin.New()
		assert.Equal(t, "MultiPlugin", mp.Name())
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		mp := multiplugin.New()
		err := mp.Process(ctx, &model.GherkinDocument{})
		assert.NoError(t, err)
	})

	t.Run("goplugin", func(t *testing.T) {
		t.Parallel()

		goPlugin := goplugin.New()

		mp := multiplugin.New(goPlugin, goPlugin)
		if assert.NotNil(t, mp) {
			err := mp.Process(ctx, &model.GherkinDocument{})
			assert.NoError(t, err)
		}
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()

		mp := multiplugin.New(errorProcessor{})
		if assert.NotNil(t, mp) {
			err := mp.Process(ctx, &model.GherkinDocument{})
			assert.Error(t, err)
		}
	})
}

type errorProcessor struct{}

func (pp errorProcessor) Process(context.Context, *model.GherkinDocument) error {
	return semerr.Error("test error")
}

func (pp errorProcessor) Name() string {
	return "errorProcessor"
}
