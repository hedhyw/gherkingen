package main

import (
	"os"

	"github.com/hedhyw/gherkingen/internal/app"
)

func main() {
	if err := app.Run(os.Args[1:], os.Stdout); err != nil {
		// nolint: forbidigo // Command-line-tool.
		println(err.Error())
		os.Exit(1)
	}
}
