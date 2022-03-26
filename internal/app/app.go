package app

import (
	"flag"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/hedhyw/gherkingen/internal/model"

	"github.com/google/uuid"
)

const (
	internalPathPrefix = "@/"
	defaultTemplate    = "std.struct.v1.go.tmpl"
)

// Run the application.
func Run(arguments []string, out io.Writer) (err error) {
	flag.CommandLine.Init(flag.CommandLine.Name(), flag.ContinueOnError)
	flag.CommandLine.SetOutput(out)

	outputFormat := flag.String(
		"format",
		string(model.FormatGo),
		"output format: "+strings.Join(model.Formats(), ", "),
	)
	templateFile := flag.String(
		"template",
		internalPathPrefix+defaultTemplate,
		"template file",
	)
	permanentIDs := flag.Bool(
		"permanent-ids",
		false,
		"The same calls to the generator always produces the same output",
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
	packageName := flag.String(
		"package",
		"generated_test",
		"name of the generated package",
	)
	if err = flag.CommandLine.Parse(arguments); err != nil {
		return err
	}

	if *permanentIDs {
		// nolint:gosec // Usage for uniq ids.
		uuid.SetRand(rand.New(rand.NewSource(0)))
	} else {
		// nolint:gosec // Usage for uniq ids.
		uuid.SetRand(rand.New(rand.NewSource(time.Now().UnixNano())))
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
		return runGenerator(
			out,
			model.Format(*outputFormat),
			*templateFile,
			inputFile,
			*packageName,
		)
	}
}
