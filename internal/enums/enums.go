package enums

type Format string

const (
	FormatJSON Format = "json"
	FormatGo   Format = "go"
	FormatRaw  Format = "raw"
)

func Formats() []string {
	return []string{
		string(FormatJSON),
		string(FormatGo),
		string(FormatRaw),
	}
}
