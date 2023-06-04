package app

import (
	"flag"
	"fmt"
	"os"
)

func runHelp(flagSet *flag.FlagSet) (err error) {
	fmt.Fprintf(flagSet.Output(), "Usage of %s [FILE]:\n", os.Args[0])
	flagSet.PrintDefaults()
	fmt.Fprintf(flagSet.Output(), "\nExample: %s example.feature >> example_test.go\n", os.Args[0])

	return nil
}
