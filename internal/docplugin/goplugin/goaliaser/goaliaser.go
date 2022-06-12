package goaliaser

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"github.com/iancoleman/strcase"
)

// Aliaser helps to convert document fields.
type Aliaser struct {
	postGoAliasReplacer *strings.Replacer
	// OK alias will be replaced consider this regexp.
	okAliasRegexp *regexp.Regexp
}

// New creates a new Aliaser for golang.
func New() *Aliaser {
	return &Aliaser{
		postGoAliasReplacer: getDefaultPostGoAliasReplacer(),
		okAliasRegexp:       regexp.MustCompile(`(Ok[A-Z]|Ok$)`),
	}
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

// StringValue returns a quoted representation of a string.
// Example: Hello "World" -> "Hello \"World\"".
func (p Aliaser) StringValue(in string) string {
	return fmt.Sprintf("%q", in)
}

// NameAlias converts a string to a valid Go capitalized alias.
// Example: hello world -> HelloWorld.
func (p Aliaser) NameAlias(in string) string {
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
