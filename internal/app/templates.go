package app

import (
	"fmt"
	"io"

	"github.com/hedhyw/gherkingen/v2/internal/assets"
)

func runListTemplates(out io.Writer) (err error) {
	templates, err := assets.Templates()
	if err != nil {
		return err
	}

	for _, t := range templates {
		fmt.Fprintln(out, internalPathPrefix+t)
	}

	return nil
}
