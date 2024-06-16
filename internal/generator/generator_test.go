package generator_test

import (
	_ "embed"
	"testing"

	"github.com/hedhyw/gherkingen/v4/internal/generator"
	"github.com/hedhyw/gherkingen/v4/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	//go:embed generator_test.en.feature
	exampleFeatureEnglish []byte

	//go:embed generator_test.en-lol.feature
	exampleFeatureLOLCAT []byte
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Language string
	}{{
		Language: "English",
	}, {
		Language: "LOLCAT",
	}}

	for _, tc := range testCases {
		t.Run(tc.Language, func(t *testing.T) {
			t.Parallel()

			args := generator.Args{
				Format:         model.FormatGo,
				InputSource:    scenarioIn(t, tc.Language),
				TemplateSource: []byte(``),
				PackageName:    "generated_test",
				Plugin:         requireNewPlugin(t),
				GenerateUUID:   uuid.NewString,
				Language:       language(t, tc.Language),
			}

			_, err := generator.Generate(args)

			assert.NoError(t, err)
		})
	}
}

func TestGenerateFailed(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Reason  string
		Prepare func(args *generator.Args)
	}{{
		Reason: "invalid_format",
		Prepare: func(args *generator.Args) {
			args.Format = "invalid"
		},
	}, {
		Reason: "invalid_source",
		Prepare: func(args *generator.Args) {
			args.InputSource = []byte("INVALID")
		},
	}, {
		Reason: "invalid_template",
		Prepare: func(args *generator.Args) {
			args.TemplateSource = []byte(`{{ .Unknown }}`)
		},
	}, {
		Reason: "no_package",
		Prepare: func(args *generator.Args) {
			args.PackageName = ""
		},
	}, {
		Reason: "unsupported_language",
		Prepare: func(args *generator.Args) {
			args.Language = "unsupported"
		},
	}}

	for _, testCase := range testCases {
		t.Run(testCase.Reason, func(t *testing.T) {
			t.Parallel()

			args := generator.Args{
				Format:         model.FormatGo,
				InputSource:    scenarioIn(t, defaultLanguage),
				TemplateSource: []byte(``),
				PackageName:    "generated_test",
				Plugin:         requireNewPlugin(t),
				GenerateUUID:   uuid.NewString,
			}

			testCase.Prepare(&args)

			_, err := generator.Generate(args)

			assert.Error(t, err)
		})
	}
}

const defaultLanguage = "English"

func scenarioIn(tb testing.TB, language string) []byte {
	tb.Helper()

	switch language {
	case "English":
		return exampleFeatureEnglish
	case "LOLCAT":
		return exampleFeatureLOLCAT
	default:
		tb.Fatalf("unexpected language: %s", language)

		return nil
	}
}

func language(tb testing.TB, name string) string {
	tb.Helper()

	switch name {
	case "English":
		return "en"
	case "LOLCAT":
		return "en-lol"
	default:
		tb.Fatalf("unexpected language name: %s", name)

		return ""
	}
}
