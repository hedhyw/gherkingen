package app

import (
	"fmt"
	"io"
)

func runVersion(w io.Writer, version string) error {
	fmt.Fprintln(w, "github.com/hedhyw/gherkingen/v2@"+version)

	return nil
}
