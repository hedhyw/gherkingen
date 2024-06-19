package generator

import (
	"fmt"
	"go/format"

	"github.com/hedhyw/gherkingen/v4/internal/model"
)

// generateGo generates raw output and formates it with go formatter.
func generateGo(
	tmplSource []byte,
	tmplData *model.TemplateData,
) (out []byte, err error) {
	out, err = generateRaw(tmplSource, tmplData)
	if err != nil {
		return nil, fmt.Errorf("generating raw: %w", err)
	}

	if out, err = format.Source(out); err != nil {
		return nil, fmt.Errorf("formating source: %w", err)
	}

	return out, nil
}
