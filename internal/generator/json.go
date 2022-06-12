package generator

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hedhyw/gherkingen/v2/internal/model"
)

// generateJSON generates output in JSON.
func generateJSON(tmplData *model.TemplateData) (data []byte, err error) {
	var buf bytes.Buffer

	jsonEncoder := json.NewEncoder(&buf)
	jsonEncoder.SetIndent("", "    ")
	if err = jsonEncoder.Encode(tmplData); err != nil {
		return nil, fmt.Errorf("encoding json: %w", err)
	}

	return buf.Bytes(), nil
}
