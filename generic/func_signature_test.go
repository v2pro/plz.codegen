package generic

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_parse_signature(t *testing.T) {
	tests := []struct {
		input    string
		expected funcSignature
	}{
		{" compare (a T, b T) int", funcSignature{
			funcName: "compare",
			funcParams: []funcParam{
				{"a", "T"},
				{"b", "T"},
			},
			funcReturns: []funcReturn{
				{"", "int"},
			},
		}},
		{" compare (a T, b T)", funcSignature{
			funcName: "compare",
			funcParams: []funcParam{
				{"a", "T"},
				{"b", "T"},
			},
			funcReturns: []funcReturn{
			},
		}},
		{" compare ()", funcSignature{
			funcName: "compare",
			funcParams: []funcParam{
			},
			funcReturns: []funcReturn{
			},
		}},
		{" compare () (a int, b err)", funcSignature{
			funcName: "compare",
			funcParams: []funcParam{
			},
			funcReturns: []funcReturn{
				{"a", "int"},
				{"b", "err"},
			},
		}},
	}
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			should := require.New(t)
			signature, err := parseSignature(test.input)
			should.Nil(err)
			should.Equal(test.expected, *signature)
		})
	}
}
