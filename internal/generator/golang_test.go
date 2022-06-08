package generator_test

import (
	"bytes"
	"go/format"
	"io"
	"testing"

	"github.com/hedhyw/gherkingen/internal/assets"
	"github.com/hedhyw/gherkingen/internal/generator"
	"github.com/hedhyw/gherkingen/internal/model"
)

func TestGenerateGo(t *testing.T) {
	t.Parallel()

	const exampleTemplate = `func Test{{upperAlias .Feature.Name}}(){}`

	gotDataGo, err := generator.Generate(model.GenerateArgs{
		Format:         model.FormatGo,
		InputSource:    exampleFeature,
		TemplateSource: []byte(exampleTemplate),
		PackageName:    "generated_test",
	})
	if err != nil {
		t.Fatal(err)
	}

	const expDataGoRaw = `func TestGuessTheWord(){}`

	expDataGo, err := format.Source([]byte(expDataGoRaw))
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(expDataGo, gotDataGo) {
		t.Fatalf("%s", gotDataGo)
	}
}

func TestGenerateGoFormattingFailed(t *testing.T) {
	t.Parallel()

	_, err := generator.Generate(model.GenerateArgs{
		Format:         model.FormatGo,
		InputSource:    exampleFeature,
		TemplateSource: []byte("-"),
		PackageName:    "generated_test",
	})
	if err == nil {
		t.Fatal(err)
	}
}

func TestGenerateAssetTemplatesShouldNotFail(t *testing.T) {
	t.Parallel()

	templates, err := assets.Templates()
	if err != nil {
		t.Fatal(err)
	}

	for _, tmpl := range templates {
		tmpl := tmpl

		t.Run(tmpl, func(t *testing.T) {
			t.Parallel()

			tmplFile, err := assets.OpenTemplate(tmpl)
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				if err := tmplFile.Close(); err != nil {
					t.Error(err)
				}
			})

			tmplData, err := io.ReadAll(tmplFile)
			if err != nil {
				t.Fatal(err)
			}

			_, err = generator.Generate(model.GenerateArgs{
				Format:         model.FormatGo,
				InputSource:    exampleFeature,
				TemplateSource: tmplData,
				PackageName:    "generated_test",
			})
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
