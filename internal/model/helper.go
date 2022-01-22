package model

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

func goString(in string) string {
	return fmt.Sprintf("%q", in)
}

func goName(in string) string {
	alias := make([]rune, 0, len(in))

	for _, r := range in {
		switch {
		case
			unicode.IsDigit(r) && len(alias) != 0,
			unicode.IsLetter(r):

			alias = append(alias, r)
		case unicode.IsSpace(r), r == '_':
			alias = append(alias, '_')
		}
	}

	if len(alias) == 0 {
		return "Undefined"
	}

	return strcase.ToCamel(string(alias))
}

func determinateGoType(values []string) goType {
	priority := [...]goType{goTypeInt, goTypeBool, goTypeFloat64, goTypeString}

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

func goValue(val string, goType goType) (string, goType) {
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

	return goString(val), goTypeString
}
