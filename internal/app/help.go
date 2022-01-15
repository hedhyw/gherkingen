package app

import (
	"flag"
	"fmt"
	"os"
)

func runHelp() (err error) {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s [FILE]:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintf(flag.CommandLine.Output(), "\nExample: %s example.feature >> example_test.go\n", os.Args[0])

	return nil
}
