package generator_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v2/internal/generator"
	"github.com/hedhyw/gherkingen/v2/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRaw(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Template string
		Exp      string
	}{
		// UpperAlias test cases.
		{
			Template: `{{upperAlias .Feature.Name}}`,
			Exp:      `GuessTheWord`,
		},
		// LowerAlias test cases.
		{
			Template: `{{lowerAlias .Feature.Name}}`,
			Exp:      `guessTheWord`,
		},
		{
			Template: `{{lowerAlias "123"}}-{{lowerAlias "123"}}-{{upperAlias "123"}}`,
			Exp:      `var1-var2-Var1`,
		},
		// TrimSpace test cases.
		{
			Template: `{{trimSpace " 123 456 "}}`,
			Exp:      `123 456`,
		},
		// PrepareGoStr test cases.
		{
			Template: `"{{prepareGoStr "test \" str"}}"`,
			Exp:      `"test \" str"`,
		},
		// WithFinalDot test cases.
		{
			Template: `{{withFinalDot "123"}}`,
			Exp:      "123.",
		},
		{
			Template: `{{withFinalDot "abc"}}`,
			Exp:      "abc.",
		},
		{
			Template: `{{withFinalDot "ABC"}}`,
			Exp:      "ABC.",
		},
		{
			Template: `{{withFinalDot "你好"}}`,
			Exp:      "你好.",
		},
		{
			Template: `{{withFinalDot "dot at the end."}}`,
			Exp:      "dot at the end.",
		},
		{
			Template: `{{withFinalDot "comma at the end,"}}`,
			Exp:      "comma at the end,",
		},
		{
			Template: `{{withFinalDot "exclamation point at the end!"}}`,
			Exp:      "exclamation point at the end!",
		},
		{
			Template: `{{withFinalDot "question mark at the end?"}}`,
			Exp:      "question mark at the end?",
		},
		{
			Template: `{{withFinalDot "😂"}}`,
			Exp:      "😂.",
		},
		{
			Template: `{{withFinalDot "\"double-quotes\""}}`,
			Exp:      "\"double-quotes\".",
		},
		{
			Template: `{{withFinalDot "'single-quotes'"}}`,
			Exp:      "'single-quotes'.",
		},
		{
			Template: `{{withFinalDot "` + string([]byte{1}) + `"}}`,
			Exp:      string([]byte{1}),
		},
		{
			Template: `{{withFinalDot "` + "`" + `code-quotes` + "`" + `"}}`,
			Exp:      "`code-quotes`.",
		},
		{
			Template: `{{withFinalDot "  "}}`,
			Exp:      "  ",
		},
		{
			Template: `{{withFinalDot "123 \n\t "}}`,
			Exp:      "123. \n\t ",
		},
		{
			Template: `{{withFinalDot ""}}`,
			Exp:      "",
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Template, func(t *testing.T) {
			t.Parallel()

			gotDataRaw, err := generator.Generate(generator.Args{
				Format:         model.FormatRaw,
				InputSource:    exampleFeature,
				TemplateSource: []byte(tc.Template),
				PackageName:    "generated_test.go",
				Plugin:         requireNewPlugin(t),
			})
			if assert.NoError(t, err) {
				assert.Equal(t, tc.Exp, string(gotDataRaw))
			}
		})
	}
}

func TestGenerateRawFailed(t *testing.T) {
	t.Parallel()

	_, err := generator.Generate(generator.Args{
		Format:         model.FormatRaw,
		InputSource:    exampleFeature,
		TemplateSource: []byte("{{"),
		PackageName:    "generated_test.go",
		Plugin:         requireNewPlugin(t),
	})
	assert.Error(t, err)
}
