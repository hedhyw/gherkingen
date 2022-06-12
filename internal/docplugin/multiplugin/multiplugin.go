package multiplugin

import (
	"context"
	"fmt"

	"github.com/hedhyw/gherkingen/internal/docplugin"
	"github.com/hedhyw/gherkingen/internal/model"
)

// MultiPlugin helps to process many plugins.
type MultiPlugin []docplugin.Plugin

// New creates a new plugin.
func New(plugins ...docplugin.Plugin) MultiPlugin {
	return MultiPlugin(plugins)
}

// Process document by all plugins.
func (pp MultiPlugin) Process(
	ctx context.Context,
	document *model.GherkinDocument,
) (err error) {
	for _, p := range pp {
		err = p.Process(ctx, document)
		if err != nil {
			return fmt.Errorf("processing: %s: %w", p.Name(), err)
		}
	}

	return nil
}

// Name of the plugin.
func (pp MultiPlugin) Name() string {
	return "MultiPlugin"
}
