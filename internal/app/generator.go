package app

import (
	"fmt"
	"io"
	"path"
	"strings"

	"github.com/hedhyw/gherkingen/v3/internal/docplugin/goplugin"
	"github.com/hedhyw/gherkingen/v3/internal/docplugin/multiplugin"
	"github.com/hedhyw/gherkingen/v3/internal/generator"
	"github.com/hedhyw/gherkingen/v3/internal/model"
)

// appArgs contains required arguments for runGenerator.
type appArgs struct {
	Output       io.Writer
	OutputFormat model.Format
	TemplateFile string
	InputFile    string
	PackageName  string
	GoParallel   bool
	GenerateUUID func() string
}

func runGenerator(
	args appArgs,
) (err error) {
	templateSource, err := readTemplate(args.TemplateFile)
	if err != nil {
		return err
	}

	if args.OutputFormat == model.FormatAutoDetect {
		args.OutputFormat = detectFormat(args.TemplateFile)
	}

	inputSource, err := readInput(args.InputFile)
	if err != nil {
		return err
	}

	goPlugin := goplugin.New(goplugin.Args{
		Parallel: args.GoParallel,
	})

	data, err := generator.Generate(generator.Args{
		Format:         args.OutputFormat,
		InputSource:    inputSource,
		TemplateSource: templateSource,
		PackageName:    args.PackageName,
		Plugin:         multiplugin.New(goPlugin),
		GenerateUUID:   args.GenerateUUID,
	})
	if err != nil {
		return err
	}

	fmt.Fprint(args.Output, string(data))

	return nil
}

func detectFormat(templateFile string) model.Format {
	templateFile = strings.ToLower(templateFile)
	switch path.Ext(templateFile) {
	case ".go":
		return model.FormatGo
	case ".tmpl":
		return detectFormat(strings.TrimSuffix(templateFile, ".tmpl"))
	default:
		return model.FormatRaw
	}
}
