package token_test

import (
	"testing"

	"github.com/Unquabain/chatgpc/token"
	"github.com/stretchr/testify/assert"
)

type ringTestCase struct {
	Description string
	Max         int
	Input       []string
	Expected    []string
}

func (tc ringTestCase) Test(t *testing.T) {
	subject := token.NewRing(tc.Max)
	for _, in := range tc.Input {
		subject.Push(in)
	}
	actual := subject.Segments()
	assert.ElementsMatch(t, tc.Expected, actual)
}

var ringTestCases = []ringTestCase{
	{
		Description: `one segment, one word`,
		Max:         1,
		Input:       []string{`one`},
		Expected:    []string{`one`},
	},
	{
		Description: `one segment, many words`,
		Max:         1,
		Input:       []string{`one`, `two`, `three`, `four`},
		Expected:    []string{`four`},
	},
	{
		Description: `many segments, many words`,
		Max:         4,
		Input:       []string{`one`, `two`, `three`, `four`},
		Expected:    []string{`one`, `two`, `three`, `four`},
	},
	{
		Description: `many segments, more words`,
		Max:         4,
		Input:       []string{`one`, `two`, `three`, `four`, `five`, `six`},
		Expected:    []string{`three`, `four`, `five`, `six`},
	},
	{
		Description: `few segments, more words`,
		Max:         2,
		Input:       []string{`one`, `two`, `three`, `four`, `five`, `six`},
		Expected:    []string{`five`, `six`},
	},
}

func TestRing(t *testing.T) {
	for _, tc := range ringTestCases {
		t.Run(tc.Description, tc.Test)
	}
}
