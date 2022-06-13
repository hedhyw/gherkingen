package goplugin

import (
	"strconv"
	"strings"

	"github.com/hedhyw/gherkingen/v2/internal/docplugin/goplugin/goaliaser"
)

// goType definition.
type goType string

// Possible go types.
const (
	goTypeString  goType = "string"
	goTypeInt     goType = "int"
	goTypeFloat64 goType = "float64"
	goTypeBool    goType = "bool"
)

func determinateGoType(values []string) goType {
	priority := [...]goType{
		goTypeInt,
		goTypeBool,
		goTypeFloat64,
		goTypeString,
	}

	goTypeCounters := make(map[goType]int, len(values))
	goTypeCounters[goTypeString] = len(values)

	for _, v := range values {
		if goInt(v) != "" {
			goTypeCounters[goTypeInt]++
		}

		if goBool(v) != "" {
			goTypeCounters[goTypeBool]++
		}

		if goFloat64(v) != "" {
			goTypeCounters[goTypeFloat64]++
		}
	}

	for _, goType := range priority {
		if goTypeCounters[goType] == len(values) {
			return goType
		}
	}

	return goTypeString
}

func goInt(val string) string {
	val = strings.ReplaceAll(val, " ", "")

	i, err := strconv.Atoi(val)
	if err != nil {
		return ""
	}

	return strconv.Itoa(i)
}

func goFloat64(val string) string {
	val = strings.ReplaceAll(val, " ", "")

	_, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return ""
	}

	return val
}

func goBool(val string) string {
	val = strings.ToLower(val)

	switch val {
	case "+":
		return strconv.FormatBool(true)
	case "-":
		return strconv.FormatBool(false)
	}

	b, err := strconv.ParseBool(val)
	if err != nil {
		return ""
	}

	return strconv.FormatBool(b)
}

func goValue(
	alias *goaliaser.Aliaser,
	val string,
	goType goType,
) (string, goType) {
	switch goType {
	case goTypeInt:
		if val := goInt(val); val != "" {
			return val, goTypeInt
		}
	case goTypeBool:
		if val := goBool(val); val != "" {
			return val, goTypeBool
		}
	case goTypeFloat64:
		if val := goFloat64(val); val != "" {
			return val, goTypeFloat64
		}
	case goTypeString:
	default:
	}

	return alias.StringValue(val), goTypeString
}
