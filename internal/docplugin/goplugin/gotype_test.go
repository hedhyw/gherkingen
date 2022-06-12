package goplugin

import (
	"strings"
	"testing"

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

	for i, tc := range testCases {
		i, tc := i, tc

		t.Run(strings.Join(tc.In, "_"), func(t *testing.T) {
			t.Parallel()

			got := determinateGoType(tc.In)
			assert.Equal(t, tc.Exp, got, i)
		})
	}
}
