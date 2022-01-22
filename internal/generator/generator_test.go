package generator_test

import (
	_ "embed"
	"testing"

	"github.com/hedhyw/gherkingen/internal/generator"
	"github.com/hedhyw/gherkingen/internal/model"
)

//go:embed generator_test.feature
var exampleFeature []byte

func TestGenerate_Failed(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name    string
		Prepare func(args *model.GenerateArgs)
		OK      bool
	}{{
		OK:      true,
		Name:    "ok",
		Prepare: func(args *model.GenerateArgs) {},
	}, {
		OK:   false,
		Name: "invalid_format",
		Prepare: func(args *model.GenerateArgs) {
			args.Format = "invalid"
		},
	}, {
		Name: "invalid_source",
		Prepare: func(args *model.GenerateArgs) {
			args.InputSource = []byte("INVALID")
		},
	}, {
		OK:   false,
		Name: "invalid_template",
		Prepare: func(args *model.GenerateArgs) {
			args.TemplateSource = []byte(`{{ .Unknown }}`)
		},
	}, {
		OK:   false,
		Name: "no_package",
		Prepare: func(args *model.GenerateArgs) {
			args.PackageName = ""
		},
	}}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			args := model.GenerateArgs{
				Format:         model.FormatGo,
				InputSource:    exampleFeature,
				TemplateSource: []byte(``),
				PackageName:    "generated_test",
			}

			tc.Prepare(&args)

			_, err := generator.Generate(args)
			if (err == nil) != tc.OK {
				t.Fatal(tc.OK, err)
			}
		})
	}
}
