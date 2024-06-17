package app

import (
	"fmt"
	"io"
	"reflect"

	gherkin "github.com/cucumber/gherkin/go/v28"
	"github.com/hedhyw/semerr/pkg/v1/semerr"
)

func runListFeatureLanguages(out io.Writer) error {
	languages, err := featureLanguages()
	if err != nil {
		return fmt.Errorf("getting feature languages: %w", err)
	}

	for _, lang := range languages {
		fmt.Fprintf(out, "%s\t%s\t%s\n", lang.Language, lang.Name, lang.Native)
	}

	return nil
}

func featureLanguages() ([]*gherkin.Dialect, error) {
	dialectProvider := gherkin.DialectsBuiltin()

	v := reflect.ValueOf(dialectProvider)

	if v.Kind() != reflect.Map {
		return nil, semerr.Error("parsing builtin dialects: unexpected type")
	}

	dialects := make([]*gherkin.Dialect, 0, len(v.MapKeys()))

	for _, key := range v.MapKeys() {
		dialect := dialectProvider.GetDialect(key.String())

		if dialect == nil {
			continue
		}

		dialects = append(dialects, dialect)
	}

	return dialects, nil
}
