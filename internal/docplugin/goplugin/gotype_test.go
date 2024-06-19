package goplugin

import (
	"strings"
	"testing"

	"github.com/hedhyw/gherkingen/v3/internal/docplugin/goplugin/goaliaser"

	"github.com/stretchr/testify/assert"
)

func TestDeterminateGoType(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		In  []string
		Exp goType
	}{{
		In:  []string{"a", "b", "c"},
		Exp: goTypeString,
	}, {
		In:  []string{"-2", "0", "2"},
		Exp: goTypeInt,
	}, {
		In:  []string{"-2", "a", "2"},
		Exp: goTypeString,
	}, {
		In:  []string{"true", "false", "true"},
		Exp: goTypeBool,
	}, {
		In:  []string{"true", "a", "true"},
		Exp: goTypeString,
	}, {
		In:  []string{"True", "FAlse"},
		Exp: goTypeBool,
	}, {
		In:  []string{"+", "-", "+"},
		Exp: goTypeBool,
	}, {
		In:  []string{"1.2", "-1.3", "0.0"},
		Exp: goTypeFloat64,
	}, {
		In:  []string{"1.2", "-1.3", "a"},
		Exp: goTypeString,
	}, {
		In:  []string{"1", "0"},
		Exp: goTypeInt,
	}, {
		In:  []string{""},
		Exp: goTypeString,
	}}

	for i, testCase := range testCases {
		t.Run(strings.Join(testCase.In, "_"), func(t *testing.T) {
			t.Parallel()

			got := determinateGoType(testCase.In)
			assert.Equal(t, testCase.Exp, got, i)
		})
	}
}

func TestGoValue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		In        string
		InGoType  goType
		Exp       string
		ExpGoType goType
	}{{
		In:        `Simple test`,
		InGoType:  goTypeString,
		Exp:       `"Simple test"`,
		ExpGoType: goTypeString,
	}, {
		In:        `Simple test`,
		InGoType:  goTypeInt,
		Exp:       `"Simple test"`,
		ExpGoType: goTypeString,
	}, {
		In:        `Simple test`,
		InGoType:  goTypeFloat64,
		Exp:       `"Simple test"`,
		ExpGoType: goTypeString,
	}, {
		In:        `Simple test`,
		InGoType:  goTypeBool,
		Exp:       `"Simple test"`,
		ExpGoType: goTypeString,
	}, {
		In:        `100`,
		InGoType:  goTypeInt,
		Exp:       `100`,
		ExpGoType: goTypeInt,
	}, {
		// nolint: dupword // False alarm.
		In:        `1 000 000`,
		InGoType:  goTypeInt,
		Exp:       `1000000`,
		ExpGoType: goTypeInt,
	}, {
		In:        `+`,
		InGoType:  goTypeBool,
		Exp:       `true`,
		ExpGoType: goTypeBool,
	}, {
		In:        `F`,
		InGoType:  goTypeBool,
		Exp:       `false`,
		ExpGoType: goTypeBool,
	}, {
		In:        `100.120`,
		InGoType:  goTypeFloat64,
		Exp:       `100.120`,
		ExpGoType: goTypeFloat64,
	}, {
		In:        `10 000.120`,
		InGoType:  goTypeFloat64,
		Exp:       `10000.120`,
		ExpGoType: goTypeFloat64,
	}}

	for i, testCase := range testCases {
		t.Run(string(testCase.InGoType)+"_"+testCase.In, func(t *testing.T) {
			t.Parallel()

			gotVal, gotType := goValue(goaliaser.New(), testCase.In, testCase.InGoType)
			if gotVal != testCase.Exp {
				t.Errorf("%d:\n\tin:  %s\n\texp: %s\n\tgot: %s", i, testCase.In, testCase.Exp, gotVal)
			}

			if gotType != testCase.ExpGoType {
				t.Errorf("%d:\n\tin:  %s\n\texp: %s\n\tgot: %s", i, testCase.In, testCase.ExpGoType, gotType)
			}
		})
	}
}
