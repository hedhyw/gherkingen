package app

import (
	"fmt"
	"io"
	"path"
	"strings"

	"github.com/hedhyw/gherkingen/v2/internal/docplugin/goplugin"
	"github.com/hedhyw/gherkingen/v2/internal/docplugin/multiplugin"
	"github.com/hedhyw/gherkingen/v2/internal/generator"
	"github.com/hedhyw/gherkingen/v2/internal/model"
)

func runGenerator(
	out io.Writer,
	outputFormat model.Format,
	templateFile string,
	inputFile string,
	packageName string,
) (err error) {
	templateSource, err := readTemplate(templateFile)
	if err != nil {
		return err
	}

	if outputFormat == model.FormatAutoDetect {
		outputFormat = detectFormat(templateFile)
	}

	inputSource, err := readInput(inputFile)
	if err != nil {
		return err
	}

	data, err := generator.Generate(generator.Args{
		Format:         outputFormat,
		InputSource:    inputSource,
		TemplateSource: templateSource,
		PackageName:    packageName,
		Plugin:         multiplugin.New(goplugin.New()),
	})
	if err != nil {
		return err
	}

	fmt.Fprint(out, string(data))

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
