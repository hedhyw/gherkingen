package generator_test

import (
	"encoding/json"
	"testing"

	"github.com/hedhyw/gherkingen/internal/generator"
	"github.com/hedhyw/gherkingen/internal/model"
)

func TestGenerateJSON(t *testing.T) {
	t.Parallel()

	gotDataJSON, err := generator.Generate(model.GenerateArgs{
		Format:         model.FormatJSON,
		InputSource:    exampleFeature,
		TemplateSource: nil,
		PackageName:    "example_json",
	})
	switch {
	case err != nil:
		t.Fatal(err)
	case len(gotDataJSON) == 0:
		t.Fatal("empty output")
	}

	var gotData map[string]interface{}
	if err = json.Unmarshal(gotDataJSON, &gotData); err != nil {
		t.Fatal(err)
	}
}
