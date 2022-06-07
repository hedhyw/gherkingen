package generator

import (
	"bytes"
	"fmt"

	"github.com/hedhyw/gherkingen/internal/docprocessor"
	"github.com/hedhyw/gherkingen/internal/model"

	"github.com/cucumber/gherkin-go/v19"
	"github.com/google/uuid"
	"github.com/hedhyw/semerr/pkg/v1/semerr"
)

// Generate generates output consider template from gherkin source.
func Generate(args model.GenerateArgs) (out []byte, err error) {
	gherkinDocument, err := gherkin.ParseGherkinDocument(
		bytes.NewReader(args.InputSource),
		uuid.NewString,
	)
	if err != nil {
		return nil, fmt.Errorf("parse document: %w", err)
	}

	if args.PackageName == "" {
		return nil, semerr.Error("package name should be defined")
	}

	docProcessor, err := docprocessor.NewProcessor()
	if err != nil {
		return nil, fmt.Errorf("creating document processor: %w", err)
	}

	tmplData := &model.TemplateData{
		GherkinDocument: (&model.GherkinDocument{}).From(gherkinDocument, docProcessor),
		PackageName:     args.PackageName,
	}

	switch args.Format {
	case model.FormatGo:
		out, err = generateGo(args.TemplateSource, tmplData)
	case model.FormatRaw:
		out, err = generateRaw(args.TemplateSource, tmplData)
	case model.FormatJSON:
		out, err = generateJSON(tmplData)
	default:
		err = semerr.Error("unknown format: " + string(args.Format))
	}
	if err != nil {
		return nil, err
	}

	return out, nil
}
