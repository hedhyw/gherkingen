package generator

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/cucumber/gherkin-go/v19"
	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
)

// GenerateRaw generates a raw outpur using an input data and a template.
func GenerateRaw(
	inputData []byte,
	templateData []byte,
) (data []byte, err error) {
	gherkinDocument, err := gherkin.ParseGherkinDocument(
		bytes.NewReader(inputData),
		uuid.NewString,
	)
	if err != nil {
		return nil, fmt.Errorf("parsing gherkin: %w", err)
	}

	var buf bytes.Buffer

	tmpl, err := template.New("template").
		Funcs(template.FuncMap{
			"upperAlias":   aliasPreparer(strcase.ToCamel),
			"lowerAlias":   aliasPreparer(strcase.ToLowerCamel),
			"trimSpace":    strings.TrimSpace,
			"prepareGoStr": prepareGoStr,
		}).
		Parse(string(templateData))
	if err != nil {
		return nil, fmt.Errorf("parsing template: %w", err)
	}

	if err = tmpl.Execute(&buf, gherkinDocument); err != nil {
		return nil, fmt.Errorf("executing template: %w", err)
	}

	return buf.Bytes(), nil
}

func aliasPreparer(postFormatter func(string) string) func(string) string {
	var i int

	return func(text string) string {
		alias := make([]rune, 0, len(text))

		for _, r := range text {
			switch {
			case
				unicode.IsDigit(r) && len(alias) != 0,
				unicode.IsLetter(r):

				alias = append(alias, r)
			case unicode.IsSpace(r), r == '_':
				alias = append(alias, '_')
			}
		}

		if len(alias) == 0 {
			i++

			return postFormatter("var" + strconv.Itoa(i))
		}

		return postFormatter(string(alias))
	}
}

func prepareGoStr(text string) string {
	text = strings.ReplaceAll(text, `"`, `\"`)

	return text
}
