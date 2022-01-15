package app

import (
	"fmt"
	"io"

	"github.com/hedhyw/gherkingen/internal/enums"
	"github.com/hedhyw/gherkingen/internal/generator"

	"github.com/hedhyw/semerr/pkg/v1/semerr"
)

func runGenerator(
	out io.Writer,
	outputFormat enums.Format,
	templateFile string,
	inputFile string,
) (err error) {
	templateData, err := readTemplate(templateFile)
	if err != nil {
		return err
	}

	inputData, err := readInput(inputFile)
	if err != nil {
		return err
	}

	var data []byte
	switch outputFormat {
	case enums.FormatGo:
		data, err = generator.GenerateGo(inputData, templateData)
	case enums.FormatRaw:
		data, err = generator.GenerateRaw(inputData, templateData)
	case enums.FormatJSON:
		data, err = generator.GenerateJSON(inputData)
	default:
		err = semerr.Error("unknown format: " + string(outputFormat))
	}
	if err != nil {
		return err
	}

	fmt.Fprint(out, string(data))

	return nil
}
