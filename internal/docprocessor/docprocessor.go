package docprocessor

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/hedhyw/gherkingen/internal/model"

	"github.com/iancoleman/strcase"
)

// Processor is a helper to convert document fields.
type Processor struct {
	postGoAliasReplacer *strings.Replacer
	// OK alias will be replaced consider this regexp.
	okAliasRegexp *regexp.Regexp
}

// NewProcessor creates a new Processor.
func NewProcessor() (*Processor, error) {
	okRe, err := regexp.Compile(`(Ok[A-Z]|Ok$)`)
	if err != nil {
		return nil, fmt.Errorf("compiling ok re: %w", err)
	}

	return &Processor{
		postGoAliasReplacer: getDefaultPostGoAliasReplacer(),
		okAliasRegexp:       okRe,
	}, nil
}

func getDefaultPostGoAliasReplacer() *strings.Replacer {
	return strings.NewReplacer(
		"Acl", "ACL",
		"Api", "API",
		"Ascii", "ASCII",
		"Cpu", "CPU",
		"Css", "CSS",
		"Dns", "DNS",
		"Eof", "EOF",
		"Guid", "GUID",
		"Html", "HTML",
		"Https", "HTTPS",
		"Http", "HTTP",
		"Id", "ID",
		"Ip", "IP",
		"Json", "JSON",
		"Lhs", "LHS",
		"Qps", "QPS",
		"Ram", "RAM",
		"Rhs", "RHS",
		"Rpc", "RPC",
		"Sla", "SLA",
		"Smtp", "SMTP",
		"Sql", "SQL",
		"Ssh", "SSH",
		"Tcp", "TCP",
		"Tls", "TLS",
		"Ttl", "TTL",
		"Udp", "UDP",
		"Uid", "UID",
		"Ui", "UI",
		"Uuid", "UUID",
		"Uri", "URI",
		"Url", "URL",
		"Utf8", "UTF8",
		"Vm", "VM",
		"Xml", "XML",
		"Xmpp", "XMPP",
		"Xsrf", "XSRF",
		"Xss", "XSS",
	)
}

// GoName returns a quored representation of a string.
// Example: Hello "World" -> "Hello \"World\""
func (p Processor) GoString(in string) string {
	return fmt.Sprintf("%q", in)
}

// GoName converts a string to a valid Go capitalized alias.
// Example: hello world -> HelloWorld.
func (p Processor) GoName(in string) string {
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

	name := string(alias)
	name = strcase.ToCamel(name)
	name = p.postGoAliasReplacer.Replace(name)
	name = p.okAliasRegexp.ReplaceAllStringFunc(name, func(s string) string {
		if len(s) < 2 {
			return s
		}

		return "OK" + s[2:]
	})

	return name
}

func (p Processor) determinateGoType(values []string) model.GoType {
	priority := [...]model.GoType{
		model.GoTypeInt,
		model.GoTypeBool,
		model.GoTypeFloat64,
		model.GoTypeString,
	}

	goTypeCounters := make(map[model.GoType]int, len(values))
	goTypeCounters[model.GoTypeString] = len(values)

	for _, v := range values {
		if p.goInt(v) != "" {
			goTypeCounters[model.GoTypeInt]++
		}

		if p.goBool(v) != "" {
			goTypeCounters[model.GoTypeBool]++
		}

		if p.goFloat64(v) != "" {
			goTypeCounters[model.GoTypeFloat64]++
		}
	}

	for _, goType := range priority {
		if goTypeCounters[goType] == len(values) {
			return goType
		}
	}

	return model.GoTypeString
}

func (p Processor) goInt(val string) string {
	val = strings.ReplaceAll(val, " ", "")

	i, err := strconv.Atoi(val)
	if err != nil {
		return ""
	}

	return strconv.Itoa(i)
}

func (p Processor) goFloat64(val string) string {
	val = strings.ReplaceAll(val, " ", "")

	_, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return ""
	}

	return val
}

func (p Processor) goBool(val string) string {
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

func (p Processor) goValue(val string, goType model.GoType) (string, model.GoType) {
	switch goType {
	case model.GoTypeInt:
		if val := p.goInt(val); val != "" {
			return val, model.GoTypeInt
		}
	case model.GoTypeBool:
		if val := p.goBool(val); val != "" {
			return val, model.GoTypeBool
		}
	case model.GoTypeFloat64:
		if val := p.goFloat64(val); val != "" {
			return val, model.GoTypeFloat64
		}
	case model.GoTypeString:
	default:
	}

	return p.GoString(val), model.GoTypeString
}
