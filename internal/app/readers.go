package app

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hedhyw/gherkingen/v2/internal/assets"

	"github.com/hedhyw/semerr/pkg/v1/semerr"
)

func readInput(inputFile string) (data []byte, err error) {
	var f io.ReadCloser
	if inputFile == "-" || inputFile == "" {
		f = io.NopCloser(os.Stdin)
	} else {
		f, err = os.Open(inputFile)
	}

	if err != nil {
		return nil, fmt.Errorf("opening gherkin: %w", err)
	}

	defer func() { err = semerr.NewMultiError(err, f.Close()) }()

	return io.ReadAll(f)
}

func readTemplate(template string) (data []byte, err error) {
	var f io.ReadCloser

	if strings.HasPrefix(template, internalPathPrefix) {
		template = strings.TrimPrefix(template, internalPathPrefix)
		f, err = assets.OpenTemplate(template)
	} else {
		f, err = os.Open(template)
	}

	if err != nil {
		return nil, fmt.Errorf("opening template: %w", err)
	}

	defer func() { err = semerr.NewMultiError(err, f.Close()) }()

	return io.ReadAll(f)
}
