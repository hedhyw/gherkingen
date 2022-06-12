package model

// Format of the output.
type Format string

// Possible formats.
const (
	FormatJSON Format = "json"
	FormatGo   Format = "go"
	FormatRaw  Format = "raw"
)

// Formats returns supported output formats.
func Formats() []string {
	return []string{
		string(FormatJSON),
		string(FormatGo),
		string(FormatRaw),
	}
}
