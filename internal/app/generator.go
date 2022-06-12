package app

import (
	"fmt"
	"io"

	"github.com/hedhyw/gherkingen/internal/docplugin/goplugin"
	"github.com/hedhyw/gherkingen/internal/docplugin/multiplugin"
	"github.com/hedhyw/gherkingen/internal/generator"
	"github.com/hedhyw/gherkingen/internal/model"
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
