package generator_test

import (
	"encoding/json"
	"testing"

	"github.com/hedhyw/gherkingen/internal/generator"
)

func TestGenerateJSON(t *testing.T) {
	t.Parallel()

	gotDataJSON, err := generator.GenerateJSON(exampleFeature)
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
