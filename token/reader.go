package token

import (
	"bufio"
	"fmt"
	"io"
	"unicode"

	"github.com/apex/log"
)

type runeTest func(rune) bool

var runeTests = []runeTest{
	unicode.IsSpace,
	unicode.IsPunct,
	unicode.IsDigit,
	unicode.IsLetter,
	func(r rune) bool {
		return !unicode.IsSpace(r) &&
			!unicode.IsPunct(r) &&
			!unicode.IsDigit(r) &&
			!unicode.IsLetter(r)
	},
}

func splitFunc(data []byte, atEOF bool) (int, []byte, error) {
	rarray := []rune(string(data))
	if len(rarray) == 0 {
		return 0, nil, nil
	}
	var (
		idx   int
		test  runeTest
		token []byte
	)
	for _, t := range runeTests {
		if t(rarray[idx]) {
			test = t
			break
		}
	}
	if test == nil {
		return 0, nil, fmt.Errorf(`could not classify rune %v`, rarray[idx])
	}
	for {
		idx++
		if idx == len(rarray) || !test(rarray[idx]) {
			break
		}
	}
	token = []byte(string(rarray[:idx]))
	if ldata, ltoken := len(data), len(token); ltoken > ldata {
		log.Warnf(`possible invalid unicode: read in %d bytes; returning %d bytes`, ldata, ltoken)
		return ldata, token[:ldata], nil
	}
	return len(token), token, nil
}

func NewReader(source io.Reader) *bufio.Scanner {
	s := bufio.NewScanner(source)
	s.Split(splitFunc)
	return s
}
