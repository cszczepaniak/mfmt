package scanner

import (
	"testing"

	"github.com/cszczepaniak/mfmt/token"
	"github.com/stretchr/testify/assert"
)

func TestScanFile(t *testing.T) {
	tests := []struct {
		name        string
		sourceFile  string
		expTokTypes []token.Type
	}{
		{
			name:        "test simple file",
			sourceFile:  "testdata/simple.m",
			expTokTypes: []token.Type{token.FLOAT, token.ELEM_MUL, token.INT, token.EOF},
		},
		{
			name:       "test if statement",
			sourceFile: "testdata/if.m",
			expTokTypes: []token.Type{token.IF, token.IDENT, token.EQL, token.INT,
				token.IDENT, token.LPAREN, token.RPAREN, token.SEMICOLON, token.END,
				token.EOF},
		},
		{
			name:       "test ellipsis",
			sourceFile: "testdata/ellipsis.m",
			expTokTypes: []token.Type{token.IDENT, token.ASSIGN, token.LBRACK,
				token.INT, token.COMMA, token.INT, token.COMMA, token.ELLIPSIS,
				token.INT, token.COMMA, token.INT, token.RBRACK, token.SEMICOLON,
				token.EOF},
		},
		{
			name:       "test classdef",
			sourceFile: "testdata/classdef.m",
			expTokTypes: []token.Type{token.CLASSDEF, token.IDENT, token.PROPERTIES,
				token.IDENT, token.IDENT, token.END, token.METHODS, token.FUNCTION,
				token.IDENT, token.ASSIGN, token.IDENT, token.LPAREN, token.RPAREN,
				token.END, token.END, token.END, token.EOF},
		},
		{
			name:       "test function handle",
			sourceFile: "testdata/fcn_hndl.m",
			expTokTypes: []token.Type{token.IDENT, token.ASSIGN, token.AT,
				token.LPAREN, token.IDENT, token.COMMA, token.IDENT, token.COMMA,
				token.IDENT, token.RPAREN, token.LPAREN, token.IDENT, token.ADD,
				token.IDENT, token.ADD, token.IDENT, token.RPAREN, token.SEMICOLON,
				token.EOF},
		},
	}
	for _, tc := range tests {
		s := ScanFile(tc.sourceFile)
		actTokTypes := make([]token.Type, 0)
		for _, t := range s.tokens {
			actTokTypes = append(actTokTypes, t.TokenType)
		}
		assert.Equal(t, tc.expTokTypes, actTokTypes, tc.name)
	}
}

func TestLineNumbers(t *testing.T) {
	tests := []struct {
		name       string
		sourceFile string
		expLines   []int
	}{
		{
			name:       "test simple file",
			sourceFile: "testdata/simple.m",
			expLines:   []int{1, 1, 1, 1},
		},
		{
			name:       "test if statement",
			sourceFile: "testdata/if.m",
			expLines:   []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3},
		},
		{
			name:       "test ellipsis",
			sourceFile: "testdata/ellipsis.m",
			expLines:   []int{1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
		},
		{
			name:       "test classdef",
			sourceFile: "testdata/classdef.m",
			expLines:   []int{1, 1, 2, 3, 4, 5, 6, 7, 7, 7, 7, 7, 7, 8, 9, 10, 10},
		},
		{
			name:       "test function handle",
			sourceFile: "testdata/fcn_hndl.m",
			expLines:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	}
	for _, tc := range tests {
		s := ScanFile(tc.sourceFile)
		actLines := make([]int, 0)
		for _, t := range s.tokens {
			actLines = append(actLines, t.Line)
		}
		assert.Equal(t, tc.expLines, actLines, tc.name)
	}
}
