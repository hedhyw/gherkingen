package assets_test

import (
	"io/ioutil"
	"testing"

	"github.com/hedhyw/gherkingen/internal/assets"
)

func TestTemplates(t *testing.T) {
	t.Parallel()

	names, err := assets.Templates()
	switch {
	case err != nil:
		t.Fatal(err)
	case len(names) == 0:
		t.Fatal(names)
	}
}

func TestOpenTemplate(t *testing.T) {
	t.Parallel()

	files := [...]string{
		"std.args.go.tmpl",
		"std.struct.go.tmpl",
	}

	for _, f := range files {
		f := f

		t.Run(f, func(t *testing.T) {
			t.Parallel()

			f, err := assets.OpenTemplate(f)
			if err != nil {
				t.Fatal(err)
			}

			defer func() {
				err := f.Close()
				if err != nil {
					t.Fatal(err)
				}
			}()

			data, err := ioutil.ReadAll(f)
			switch {
			case err != nil:
				t.Fatal(err)
			case len(data) == 0:
				t.Fatal()
			}
		})
	}
}
