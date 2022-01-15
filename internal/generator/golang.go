package generator

import (
	"fmt"
	"go/format"
)

// GenerateJSON generates raw output and formates it with go formatter.
func GenerateGo(
	inputData []byte,
	templateData []byte,
) (data []byte, err error) {
	data, err = GenerateRaw(inputData, templateData)
	if err != nil {
		return nil, fmt.Errorf("generating raw: %w", err)
	}

	if data, err = format.Source(data); err != nil {
		return nil, fmt.Errorf("formating source: %w", err)
	}

	return data, nil
}
