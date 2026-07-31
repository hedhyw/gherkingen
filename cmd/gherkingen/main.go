package main

import (
	"os"

	"github.com/hedhyw/gherkingen/v4/internal/app"
)

// Version will be set on build.
var version = "unknown"

func main() {
	err := app.Run(os.Args[1:], os.Stdout, version)
	if err != nil {
		// nolint: forbidigo // Command-line-tool.
		println(err.Error())
		os.Exit(1)
	}
}
