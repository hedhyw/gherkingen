package assets

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed *.tmpl
var templates embed.FS

// OpenTemplate opens the embed template by name.
func OpenTemplate(name string) (fs.File, error) {
	return templates.Open(name)
}

// Templates returns a list of embed templates.
func Templates() ([]string, error) {
	files, err := templates.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("reading dir: %w", err)
	}

	fileNames := make([]string, 0, len(files))
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}

	return fileNames, nil
}
