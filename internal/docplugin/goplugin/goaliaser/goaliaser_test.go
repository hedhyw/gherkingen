package goaliaser_test

import (
	"testing"

	"github.com/hedhyw/gherkingen/v3/internal/docplugin/goplugin/goaliaser"

	"github.com/stretchr/testify/assert"
)

func TestStringValue(t *testing.T) {
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

	p := goaliaser.New()

	for i, tc := range testCases {
		i, tc := i, tc

		t.Run(tc.In, func(t *testing.T) {
			t.Parallel()

			got := p.StringValue(tc.In)
			assert.Equal(t, tc.Exp, got, i)
		})
	}
}

func TestNameAlias(t *testing.T) {
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

	p := goaliaser.New()

	for i, tc := range testCases {
		i, tc := i, tc

		t.Run(tc.In, func(t *testing.T) {
			t.Parallel()

			got := p.NameAlias(tc.In)
			assert.Equal(t, tc.Exp, got, i)
		})
	}
}
