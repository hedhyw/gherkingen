package generator

import (
	"bytes"
	"context"
	"fmt"

	"github.com/hedhyw/gherkingen/v4/internal/docplugin"
	"github.com/hedhyw/gherkingen/v4/internal/model"

	gherkin "github.com/cucumber/gherkin/go/v28"
	"github.com/hedhyw/semerr/pkg/v1/semerr"
)

// Args contains required arguments for Generate.
type Args struct {
	Format         model.Format
	InputSource    []byte
	TemplateSource []byte
	PackageName    string
	Plugin         docplugin.Plugin
	GenerateUUID   func() string
}

// Generate generates output consider template from gherkin source.
func Generate(args Args) (out []byte, err error) {
	gherkinDocument, err := gherkin.ParseGherkinDocument(
		bytes.NewReader(args.InputSource),
		args.GenerateUUID,
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
