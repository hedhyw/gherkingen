package enums_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/internal/enums"
)

func TestFormats(t *testing.T) {
	t.Parallel()

	formats := enums.Formats()

	expFormats := [...]enums.Format{
		enums.FormatJSON,
		enums.FormatGo,
		enums.FormatRaw,
	}

	formatsSet := make(map[string]struct{}, len(formats))
	for _, f := range formats {
		formatsSet[f] = struct{}{}
	}

	if len(formats) != len(expFormats) {
		t.Fatal(len(formats), formats)
	}

	for _, f := range expFormats {
		_, ok := formatsSet[string(f)]
		if !ok {
			t.Errorf("format %s not found in %s", f, formats)
		}
	}
}
