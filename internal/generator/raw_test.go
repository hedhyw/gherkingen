package generator_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/internal/generator"
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

			gotDataRaw, err := generator.GenerateRaw(
				exampleFeature,
				[]byte(tc.Template),
			)
			if err != nil {
				t.Fatal(err)
			}

			if tc.Exp != string(gotDataRaw) {
				t.Fatalf("%s", gotDataRaw)
			}
		})
	}
}
