package generator_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/internal/generator"
	"github.com/hedhyw/gherkingen/internal/model"
)

func TestGenerateRaw(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Template string
		Exp      string
	}{{
		Template: `{{upperAlias .Feature.Name}}`,
		Exp:      `GuessTheWord`,
	}, {
		Template: `{{lowerAlias .Feature.Name}}`,
		Exp:      `guessTheWord`,
	}, {
		Template: `{{lowerAlias "123"}}-{{lowerAlias "123"}}-{{upperAlias "123"}}`,
		Exp:      `var1-var2-Var1`,
	}, {
		Template: `{{trimSpace " 123 456 "}}`,
		Exp:      `123 456`,
	}, {
		Template: `"{{prepareGoStr "test \" str"}}"`,
		Exp:      `"test \" str"`,
	}}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Template, func(t *testing.T) {
			t.Parallel()

			gotDataRaw, err := generator.Generate(model.GenerateArgs{
				Format:         model.FormatRaw,
				InputSource:    exampleFeature,
				TemplateSource: []byte(tc.Template),
				PackageName:    "generated_test.go",
			})
			if err != nil {
				t.Fatal(err)
			}

			if tc.Exp != string(gotDataRaw) {
				t.Fatalf("%s", gotDataRaw)
			}
		})
	}
}

func TestGenerateRaw_failed(t *testing.T) {
	t.Parallel()

	_, err := generator.Generate(model.GenerateArgs{
		Format:         model.FormatRaw,
		InputSource:    exampleFeature,
		TemplateSource: []byte("{{"),
		PackageName:    "generated_test.go",
	})
	if err == nil {
		t.Fatal(err)
	}
}
