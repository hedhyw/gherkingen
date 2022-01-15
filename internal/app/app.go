package app

import (
	"flag"
	"io"
	"strings"

	"github.com/hedhyw/gherkingen/internal/enums"
)

const (
	internalPathPrefix = "@/"
	defaultTemplate    = "std.struct.go.tmpl"
)

// Run the application.
func Run(arguments []string, out io.Writer) (err error) {
	flag.CommandLine.SetOutput(out)

	outputFormat := flag.String(
		"format",
		string(enums.FormatGo),
		"output format: "+strings.Join(enums.Formats(), ", "),
	)
	templateFile := flag.String(
		"template",
		internalPathPrefix+defaultTemplate,
		"template file",
	)
	help := flag.Bool(
		"help",
		false,
		"print usage",
	)
	list := flag.Bool(
		"list",
		false,
		"list internal templates",
	)
	if err = flag.CommandLine.Parse(arguments); err != nil {
		return err
	}

	var inputFile string
	if flag.NArg() == 1 {
		inputFile = flag.Args()[0]
	}

	switch {
	case *list:
		return runListTemplates(out)
	case *help, inputFile == "":
		return runHelp()
	default:
		return runGenerator(out, enums.Format(*outputFormat), *templateFile, inputFile)
	}
}
