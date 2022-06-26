package generator

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hedhyw/gherkingen/v2/internal/docplugin"
	"github.com/hedhyw/gherkingen/v2/internal/model"

	gherkin "github.com/cucumber/common/gherkin/go/v24"
	"github.com/google/uuid"
	"github.com/hedhyw/semerr/pkg/v1/semerr"
)

// Args contains required arguments for Generate.
type Args struct {
	Format         model.Format
	InputSource    []byte
	TemplateSource []byte
	PackageName    string
	Plugin         docplugin.Plugin
}

// Generate generates output consider template from gherkin source.
func Generate(args Args) (out []byte, err error) {
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

	doc := (&model.GherkinDocument{}).From(gherkinDocument)

	err = args.Plugin.Process(context.Background(), doc)
	if err != nil {
		return nil, fmt.Errorf("processing plugin: %s: %w", args.Plugin.Name(), err)
	}

	tmplData := &model.TemplateData{
		GherkinDocument: doc,
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
