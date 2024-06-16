package generator_test

import (
	"encoding/json"
	"testing"

	"github.com/hedhyw/gherkingen/v4/internal/generator"
	"github.com/hedhyw/gherkingen/v4/internal/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJSON(t *testing.T) {
	t.Parallel()

	gotDataJSON, err := generator.Generate(generator.Args{
		Format:         model.FormatJSON,
		InputSource:    exampleFeatureEnglish,
		TemplateSource: nil,
		PackageName:    "example_json",
		Plugin:         requireNewPlugin(t),
		GenerateUUID:   uuid.NewString,
	})
	if assert.NoError(t, err) {
		assert.NotEmpty(t, gotDataJSON)

		var gotData map[string]any
		err = json.Unmarshal(gotDataJSON, &gotData)
		assert.NoError(t, err)
	}
}
