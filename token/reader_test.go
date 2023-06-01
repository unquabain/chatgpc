package token_test

import (
	"bytes"
	"testing"

	"github.com/Unquabain/chatgpc/token"
	"github.com/stretchr/testify/assert"
)

type readerTestCase struct {
	Description string
	Input       []byte
	Expected    []string
}

func (tc readerTestCase) Test(t *testing.T) {
	var actual []string = nil
	assert := assert.New(t)
	reader := token.NewReader(bytes.NewBuffer(tc.Input))
	for reader.Scan() {
		actual = append(actual, reader.Text())
	}
	assert.NoError(reader.Err())
	assert.ElementsMatch(tc.Expected, actual)
}

var readerTestCases = []readerTestCase{
	{
		Description: `happy path`,
		Input:       []byte(`this and that 42.`),
		Expected:    []string{`this`, ` `, `and`, ` `, `that`, ` `, `42`, `.`},
	},
	{
		Description: `empty string`,
		Input:       []byte(``),
		Expected:    []string{},
	},
	{
		Description: `weird text`,
		Input:       []byte(`http://localhost:2234/base64encoding`),
		Expected:    []string{`http`, `://`, `localhost`, `:`, `2234`, `/`, `base`, `64`, `encoding`},
	},
	{
		Description: `control characters`,
		Input:       []byte("this\tthat\nthis word is \033[1mbold\033[0m"),
		Expected:    []string{`this`, "\t", `that`, "\n", `this`, ` `, `word`, ` `, `is`, ` `, "\033", `[`, `1`, `mbold`, "\033", `[`, `0`, `m`},
	},
}

func TestReader(t *testing.T) {
	for _, tc := range readerTestCases {
		t.Run(tc.Description, tc.Test)
	}
}
