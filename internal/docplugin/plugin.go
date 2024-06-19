package docplugin

import (
	"context"

	"github.com/hedhyw/gherkingen/v4/internal/model"
)

// Plugin injects extra data to document.
type Plugin interface {
	// Process document.
	Process(ctx context.Context, document *model.GherkinDocument) (err error)
	// Name of the plugin.
	Name() string
}
