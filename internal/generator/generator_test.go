package generator_test

import (
	_ "embed"
	"testing"

	"github.com/hedhyw/gherkingen/v4/internal/generator"
	"github.com/hedhyw/gherkingen/v4/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

//go:embed generator_test.feature
var exampleFeature []byte

func TestGenerateFailed(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name    string
		Prepare func(args *generator.Args)
		OK      bool
	}{{
		OK:      true,
		Name:    "ok",
		Prepare: func(*generator.Args) {},
	}, {
		OK:   false,
		Name: "invalid_format",
		Prepare: func(args *generator.Args) {
			args.Format = "invalid"
		},
	}, {
		Name: "invalid_source",
		Prepare: func(args *generator.Args) {
			args.InputSource = []byte("INVALID")
		},
	}, {
		OK:   false,
		Name: "invalid_template",
		Prepare: func(args *generator.Args) {
			args.TemplateSource = []byte(`{{ .Unknown }}`)
		},
	}, {
		OK:   false,
		Name: "no_package",
		Prepare: func(args *generator.Args) {
			args.PackageName = ""
		},
	}}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			args := generator.Args{
				Format:         model.FormatGo,
				InputSource:    exampleFeature,
				TemplateSource: []byte(``),
				PackageName:    "generated_test",
				Plugin:         requireNewPlugin(t),
				GenerateUUID:   uuid.NewString,
			}

			testCase.Prepare(&args)

			_, err := generator.Generate(args)
			if testCase.OK {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
