package app

import (
	"flag"
	"io"
	"math/rand"
	"strings"
	"time"

	"github.com/hedhyw/gherkingen/v4/internal/model"

	"github.com/google/uuid"
)

const (
	internalPathPrefix       = "@/"
	defaultTemplate          = "std.simple.v1.go.tmpl"
	defaultDisableGoParallel = false
	defaultOutputFormat      = model.FormatAutoDetect
)

// Run the application.
func Run(arguments []string, out io.Writer, version string) (err error) {
	flagSet := flag.NewFlagSet(flag.CommandLine.Name(), flag.ContinueOnError)
	flagSet.SetOutput(out)

	outputFormat := flagSet.String(
		"format",
		string(defaultOutputFormat),
		"output format: "+strings.Join(model.Formats(), ", "),
	)
	templateFile := flagSet.String(
		"template",
		internalPathPrefix+defaultTemplate,
		"template file",
	)
	permanentIDs := flagSet.Bool(
		"permanent-ids",
		false,
		"The same calls to the generator always produces the same output",
	)
	helpCmd := flagSet.Bool(
		"help",
		false,
		"print usage",
	)
	_ = flagSet.Bool(
		"go-parallel",
		!defaultDisableGoParallel,
		"add parallel mark (deprecated, enabled by default)",
	)
	disableGoParallel := flagSet.Bool(
		"disable-go-parallel",
		defaultDisableGoParallel,
		"disable execution of tests in parallel",
	)
	listCmd := flagSet.Bool(
		"list",
		false,
		"list internal templates",
	)
	packageName := flagSet.String(
		"package",
		"generated_test",
		"name of the generated package",
	)
	versionCmd := flagSet.Bool(
		"version",
		false,
		"print version",
	)
	if err = flagSet.Parse(arguments); err != nil {
		return err
	}

	var seed int64

	if *permanentIDs {
		seed = 1
	} else {
		seed = time.Now().UnixNano()
	}

	var inputFile string
	if flagSet.NArg() == 1 {
		inputFile = flagSet.Args()[0]
	}

	switch {
	case *versionCmd:
		return runVersion(out, version)
	case *listCmd:
		return runListTemplates(out)
	case *helpCmd, inputFile == "":
		return runHelp(flagSet)
	default:
		return runGenerator(appArgs{
			Output:       out,
			OutputFormat: model.Format(*outputFormat),
			TemplateFile: *templateFile,
			InputFile:    inputFile,
			PackageName:  *packageName,
			GoParallel:   !(*disableGoParallel),
			GenerateUUID: newUUIDRandomGenerator(seed),
		})
	}
}

func newUUIDRandomGenerator(seed int64) func() string {
	// nolint:gosec // Usage for uniq ids.
	randomGenerator := rand.New(rand.NewSource(seed))

	return func() string {
		uuidValue, err := uuid.NewRandomFromReader(randomGenerator)

		return uuid.Must(uuidValue, err).String()
	}
}
