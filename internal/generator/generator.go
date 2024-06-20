package generator

import (
	"bytes"
	"context"
	"fmt"

	gherkin "github.com/cucumber/gherkin/go/v28"
	"github.com/hedhyw/semerr/pkg/v1/semerr"

	"github.com/hedhyw/gherkingen/v4/internal/docplugin"
	"github.com/hedhyw/gherkingen/v4/internal/model"
)

// Args contains required arguments for Generate.
type Args struct {
	Format         model.Format
	InputSource    []byte
	TemplateSource []byte
	PackageName    string
	Plugin         docplugin.Plugin
	GenerateUUID   func() string
	Language       string
}

// Generate generates output consider template from gherkin source.
func Generate(args Args) (out []byte, err error) {
	dialect, err := getDialect(args.Language)
	if err != nil {
		return nil, fmt.Errorf("getting dialect: %w", err)
	}

	gherkinDocument, err := gherkin.ParseGherkinDocumentForLanguage(
		bytes.NewReader(args.InputSource),
		dialect.Language,
		args.GenerateUUID,
	)
	if err != nil {
		return nil, fmt.Errorf("parsing document: %w", err)
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

func getDialect(language string) (*gherkin.Dialect, error) {
	if language == "" {
		language = gherkin.DefaultDialect
	}

	dialect := gherkin.DialectsBuiltin().GetDialect(language)

	if dialect == nil {
		return nil, semerr.Error("language is not supported: " + language)
	}

	return dialect, nil
}
