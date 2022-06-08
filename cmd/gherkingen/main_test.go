package main

import (
	"os"
	"testing"
)

func TestMainVersion(_ *testing.T) {
	os.Args = append(os.Args, "-version")
	main()
}
