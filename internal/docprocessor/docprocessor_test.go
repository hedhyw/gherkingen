package docprocessor

import (
	"strings"
	"testing"

	"github.com/hedhyw/gherkingen/internal/model"
)

func TestGoString(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		In  string
		Exp string
	}{{
		In:  `Simple test`,
		Exp: `"Simple test"`,
	}, {
		In:  `"Simple test"`,
		Exp: `"\"Simple test\""`,
	}, {
		In:  "`Simple test`",
		Exp: "\"`Simple test`\"",
	}, {
		In:  "`Simple \"test\"`",
		Exp: "\"`Simple \\\"test\\\"`\"",
	}, {
		In:  "xin chào",
		Exp: `"xin chào"`,
	}}

	p, err := NewProcessor()
	if err != nil {
		t.Fatal(err)
	}

	for i, tc := range testCases {
		i, tc := i, tc

		t.Run(tc.In, func(t *testing.T) {
			t.Parallel()

			got := p.GoString(tc.In)
			if got != tc.Exp {
				t.Fatalf("%d:\n\tin:  %s\n\texp: %s\n\tgot: %s", i, tc.In, tc.Exp, got)
			}
		})
	}
}

func TestGoName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		In  string
		Exp string
	}{{
		In:  "Application command line tool",
		Exp: "ApplicationCommandLineTool",
	}, {
		In:  "Application \"command\" line tool.",
		Exp: "ApplicationCommandLineTool",
	}, {
		In:  "application_command_line_tool.",
		Exp: "ApplicationCommandLineTool",
	}, {
		In:  "123app",
		Exp: "App",
	}, {
		In:  "App123",
		Exp: "App123",
	}, {
		In:  "123+.",
		Exp: "Undefined",
	}, {
		In:  "Ok",
		Exp: "OK",
	}, {
		In:  "Okay",
		Exp: "Okay",
	}, {
		In:  "Ok Go",
		Exp: "OKGo",
	}, {
		In:  "Acl",
		Exp: "ACL",
	}, {
		In:  "Api",
		Exp: "API",
	}, {
		In:  "Ascii",
		Exp: "ASCII",
	}, {
		In:  "Cpu",
		Exp: "CPU",
	}, {
		In:  "Css",
		Exp: "CSS",
	}, {
		In:  "Dns",
		Exp: "DNS",
	}, {
		In:  "Eof",
		Exp: "EOF",
	}, {
		In:  "Guid",
		Exp: "GUID",
	}, {
		In:  "Html",
		Exp: "HTML",
	}, {
		In:  "Https",
		Exp: "HTTPS",
	}, {
		In:  "Http",
		Exp: "HTTP",
	}, {
		In:  "Id",
		Exp: "ID",
	}, {
		In:  "Ip",
		Exp: "IP",
	}, {
		In:  "Json",
		Exp: "JSON",
	}, {
		In:  "Lhs",
		Exp: "LHS",
	}, {
		In:  "Qps",
		Exp: "QPS",
	}, {
		In:  "Ram",
		Exp: "RAM",
	}, {
		In:  "Rhs",
		Exp: "RHS",
	}, {
		In:  "Rpc",
		Exp: "RPC",
	}, {
		In:  "Sla",
		Exp: "SLA",
	}, {
		In:  "Smtp",
		Exp: "SMTP",
	}, {
		In:  "Sql",
		Exp: "SQL",
	}, {
		In:  "Ssh",
		Exp: "SSH",
	}, {
		In:  "Tcp",
		Exp: "TCP",
	}, {
		In:  "Tls",
		Exp: "TLS",
	}, {
		In:  "Ttl",
		Exp: "TTL",
	}, {
		In:  "Udp",
		Exp: "UDP",
	}, {
		In:  "Ui",
		Exp: "UI",
	}, {
		In:  "Uid",
		Exp: "UID",
	}, {
		In:  "Uuid",
		Exp: "UUID",
	}, {
		In:  "Uri",
		Exp: "URI",
	}, {
		In:  "Url",
		Exp: "URL",
	}, {
		In:  "Utf8",
		Exp: "UTF8",
	}, {
		In:  "Vm",
		Exp: "VM",
	}, {
		In:  "Xml",
		Exp: "XML",
	}, {
		In:  "Xmpp",
		Exp: "XMPP",
	}, {
		In:  "Xsrf",
		Exp: "XSRF",
	}, {
		In:  "Xss",
		Exp: "XSS",
	}, {
		In:  "Hello xml and json",
		Exp: "HelloXMLAndJSON",
	}}

	p, err := NewProcessor()
	if err != nil {
		t.Fatal(err)
	}

	for i, tc := range testCases {
		i, tc := i, tc

		t.Run(tc.In, func(t *testing.T) {
			t.Parallel()

			got := p.GoName(tc.In)
			if got != tc.Exp {
				t.Fatalf("%d: exp: %s got: %s", i, tc.Exp, got)
			}
		})
	}
}

func TestDeterminateGoType(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		In  []string
		Exp model.GoType
	}{{
		In:  []string{"a", "b", "c"},
		Exp: model.GoTypeString,
	}, {
		In:  []string{"-2", "0", "2"},
		Exp: model.GoTypeInt,
	}, {
		In:  []string{"-2", "a", "2"},
		Exp: model.GoTypeString,
	}, {
		In:  []string{"true", "false", "true"},
		Exp: model.GoTypeBool,
	}, {
		In:  []string{"true", "a", "true"},
		Exp: model.GoTypeString,
	}, {
		In:  []string{"True", "FAlse"},
		Exp: model.GoTypeBool,
	}, {
		In:  []string{"+", "-", "+"},
		Exp: model.GoTypeBool,
	}, {
		In:  []string{"1.2", "-1.3", "0.0"},
		Exp: model.GoTypeFloat64,
	}, {
		In:  []string{"1.2", "-1.3", "a"},
		Exp: model.GoTypeString,
	}, {
		In:  []string{"1", "0"},
		Exp: model.GoTypeInt,
	}}

	p, err := NewProcessor()
	if err != nil {
		t.Fatal(err)
	}

	for i, tc := range testCases {
		i, tc := i, tc

		t.Run(strings.Join(tc.In, "_"), func(t *testing.T) {
			t.Parallel()

			got := p.determinateGoType(tc.In)
			if got != tc.Exp {
				t.Fatalf("%d: exp: %s got: %s", i, tc.Exp, got)
			}
		})
	}
}

func TestGoValue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		In        string
		InGoType  model.GoType
		Exp       string
		ExpGoType model.GoType
	}{{
		In:        `Simple test`,
		InGoType:  model.GoTypeString,
		Exp:       `"Simple test"`,
		ExpGoType: model.GoTypeString,
	}, {
		In:        `Simple test`,
		InGoType:  model.GoTypeInt,
		Exp:       `"Simple test"`,
		ExpGoType: model.GoTypeString,
	}, {
		In:        `Simple test`,
		InGoType:  model.GoTypeFloat64,
		Exp:       `"Simple test"`,
		ExpGoType: model.GoTypeString,
	}, {
		In:        `Simple test`,
		InGoType:  model.GoTypeBool,
		Exp:       `"Simple test"`,
		ExpGoType: model.GoTypeString,
	}, {
		In:        `100`,
		InGoType:  model.GoTypeInt,
		Exp:       `100`,
		ExpGoType: model.GoTypeInt,
	}, {
		In:        `1 000 000`,
		InGoType:  model.GoTypeInt,
		Exp:       `1000000`,
		ExpGoType: model.GoTypeInt,
	}, {
		In:        `+`,
		InGoType:  model.GoTypeBool,
		Exp:       `true`,
		ExpGoType: model.GoTypeBool,
	}, {
		In:        `F`,
		InGoType:  model.GoTypeBool,
		Exp:       `false`,
		ExpGoType: model.GoTypeBool,
	}, {
		In:        `100.120`,
		InGoType:  model.GoTypeFloat64,
		Exp:       `100.120`,
		ExpGoType: model.GoTypeFloat64,
	}, {
		In:        `10 000.120`,
		InGoType:  model.GoTypeFloat64,
		Exp:       `10000.120`,
		ExpGoType: model.GoTypeFloat64,
	}}

	p, err := NewProcessor()
	if err != nil {
		t.Fatal(err)
	}

	for i, tc := range testCases {
		i, tc := i, tc

		t.Run(string(tc.InGoType)+"_"+tc.In, func(t *testing.T) {
			t.Parallel()

			gotVal, gotType := p.goValue(tc.In, tc.InGoType)
			if gotVal != tc.Exp {
				t.Errorf("%d:\n\tin:  %s\n\texp: %s\n\tgot: %s", i, tc.In, tc.Exp, gotVal)
			}

			if gotType != tc.ExpGoType {
				t.Errorf("%d:\n\tin:  %s\n\texp: %s\n\tgot: %s", i, tc.In, tc.ExpGoType, gotType)
			}
		})
	}
}
