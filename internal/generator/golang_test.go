package generator_test

import (
	"go/format"
	"io"
	"testing"

	"github.com/hedhyw/gherkingen/internal/assets"
	"github.com/hedhyw/gherkingen/internal/docplugin/goplugin"
	"github.com/hedhyw/gherkingen/internal/docplugin/multiplugin"
	"github.com/hedhyw/gherkingen/internal/generator"
	"github.com/hedhyw/gherkingen/internal/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateGo(t *testing.T) {
	t.Parallel()

	const exampleTemplate = `func Test{{upperAlias .Feature.Name}}(){}`

	gotDataGo, err := generator.Generate(generator.Args{
		Format:         model.FormatGo,
		InputSource:    exampleFeature,
		TemplateSource: []byte(exampleTemplate),
		PackageName:    "generated_test",
		Plugin:         requireNewPlugin(t),
	})
	require.NoError(t, err)

	const expDataGoRaw = `func TestGuessTheWord(){}`

	expDataGo, err := format.Source([]byte(expDataGoRaw))
	if assert.NoError(t, err) {
		assert.Equal(t, expDataGo, gotDataGo)
	}
}

func TestGenerateGoFormattingFailed(t *testing.T) {
	t.Parallel()

	_, err := generator.Generate(generator.Args{
		Format:         model.FormatGo,
		InputSource:    exampleFeature,
		TemplateSource: []byte("-"),
		PackageName:    "generated_test",
		Plugin:         requireNewPlugin(t),
	})
	assert.Error(t, err)
}

func TestGenerateAssetTemplatesShouldNotFail(t *testing.T) {
	t.Parallel()

	templates, err := assets.Templates()
	require.NoError(t, err)

	for _, tmpl := range templates {
		tmpl := tmpl

		t.Run(tmpl, func(t *testing.T) {
			t.Parallel()

			tmplFile, err := assets.OpenTemplate(tmpl)
			require.NoError(t, err)

			t.Cleanup(func() { assert.NoError(t, tmplFile.Close()) })

			tmplData, err := io.ReadAll(tmplFile)
			require.NoError(t, err)

			_, err = generator.Generate(generator.Args{
				Format:         model.FormatGo,
				InputSource:    exampleFeature,
				TemplateSource: tmplData,
				PackageName:    "generated_test",
				Plugin:         requireNewPlugin(t),
			})
			require.NoError(t, err)
		})
	}
}

func requireNewPlugin(_ testing.TB) multiplugin.MultiPlugin {
	return multiplugin.New(goplugin.New())
}
