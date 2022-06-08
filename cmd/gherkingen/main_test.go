package main

import (
	"os"
	"testing"
)

func TestMainVersion(t *testing.T) {
	os.Args = append(os.Args, "-version")
	main()
}
