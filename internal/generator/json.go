package generator

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cucumber/gherkin-go/v19"
	"github.com/google/uuid"
)

func GenerateJSON(
	inputData []byte,
) (data []byte, err error) {
	gherkinDocument, err := gherkin.ParseGherkinDocument(
		bytes.NewReader(inputData),
		uuid.NewString,
	)
	if err != nil {
		return nil, fmt.Errorf("parse document: %w", err)
	}

	var buf bytes.Buffer

	jsonEncoder := json.NewEncoder(&buf)
	jsonEncoder.SetIndent("", "    ")
	if err = jsonEncoder.Encode(gherkinDocument); err != nil {
		return nil, fmt.Errorf("encoding json: %w", err)
	}

	return buf.Bytes(), nil
}
