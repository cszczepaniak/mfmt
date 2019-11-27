package scanner

import (
	"testing"

	"github.com/cszczepaniak/mfmt/token"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanFile(t *testing.T) {
	tests := []struct {
		name       string
		sourceFile string
		expToks    []token.Token
	}{
		{
			name:       "test simple file",
			sourceFile: "testdata/simple.m",
			expToks: []token.Token{
				{TokenType: token.FLOAT, Lexeme: `.123`, Line: 1},
				{TokenType: token.ELEM_MUL, Lexeme: `.*`, Line: 1},
				{TokenType: token.INT, Lexeme: `123`, Line: 1},
				{TokenType: token.EOF, Lexeme: ``, Line: 1},
			},
		},
		{
			name:       "test if statement",
			sourceFile: "testdata/if.m",
			expToks: []token.Token{
				{TokenType: token.IF, Lexeme: `if`, Line: 1},
				{TokenType: token.IDENT, Lexeme: `myVar`, Line: 1},
				{TokenType: token.EQL, Lexeme: `==`, Line: 1},
				{TokenType: token.INT, Lexeme: `5`, Line: 1},
				{TokenType: token.IDENT, Lexeme: `doSomething`, Line: 2},
				{TokenType: token.LPAREN, Lexeme: `(`, Line: 2},
				{TokenType: token.RPAREN, Lexeme: `)`, Line: 2},
				{TokenType: token.SEMICOLON, Lexeme: `;`, Line: 2},
				{TokenType: token.END, Lexeme: `end`, Line: 3},
				{TokenType: token.EOF, Lexeme: ``, Line: 3},
			},
		},
		{
			name:       "test ellipsis",
			sourceFile: "testdata/ellipsis.m",
			expToks: []token.Token{
				{TokenType: token.IDENT, Lexeme: `A`, Line: 1},
				{TokenType: token.ASSIGN, Lexeme: `=`, Line: 1},
				{TokenType: token.LBRACK, Lexeme: `[`, Line: 1},
				{TokenType: token.INT, Lexeme: `1`, Line: 1},
				{TokenType: token.COMMA, Lexeme: `,`, Line: 1},
				{TokenType: token.INT, Lexeme: `2`, Line: 1},
				{TokenType: token.COMMA, Lexeme: `,`, Line: 1},
				{TokenType: token.ELLIPSIS, Lexeme: `...`, Line: 1},
				{TokenType: token.INT, Lexeme: `3`, Line: 2},
				{TokenType: token.COMMA, Lexeme: `,`, Line: 2},
				{TokenType: token.INT, Lexeme: `4`, Line: 2},
				{TokenType: token.RBRACK, Lexeme: `]`, Line: 2},
				{TokenType: token.SEMICOLON, Lexeme: `;`, Line: 2},
				{TokenType: token.EOF, Lexeme: ``, Line: 2},
			},
		},
		{
			name:       "test classdef",
			sourceFile: "testdata/classdef.m",
			expToks: []token.Token{
				{TokenType: token.CLASSDEF, Lexeme: `classdef`, Line: 1},
				{TokenType: token.IDENT, Lexeme: `MyClass`, Line: 1},
				{TokenType: token.PROPERTIES, Lexeme: `properties`, Line: 2},
				{TokenType: token.IDENT, Lexeme: `Abc`, Line: 3},
				{TokenType: token.IDENT, Lexeme: `Def`, Line: 4},
				{TokenType: token.END, Lexeme: `end`, Line: 5},
				{TokenType: token.METHODS, Lexeme: `methods`, Line: 6},
				{TokenType: token.FUNCTION, Lexeme: `function`, Line: 7},
				{TokenType: token.IDENT, Lexeme: `obj`, Line: 7},
				{TokenType: token.ASSIGN, Lexeme: `=`, Line: 7},
				{TokenType: token.IDENT, Lexeme: `MyClass`, Line: 7},
				{TokenType: token.LPAREN, Lexeme: `(`, Line: 7},
				{TokenType: token.RPAREN, Lexeme: `)`, Line: 7},
				{TokenType: token.END, Lexeme: `end`, Line: 8},
				{TokenType: token.END, Lexeme: `end`, Line: 9},
				{TokenType: token.END, Lexeme: `end`, Line: 10},
				{TokenType: token.EOF, Lexeme: ``, Line: 10},
			},
		},
		{
			name:       "test function handle",
			sourceFile: "testdata/fcn_hndl.m",
			expToks: []token.Token{
				{TokenType: token.IDENT, Lexeme: `fcn`, Line: 1},
				{TokenType: token.ASSIGN, Lexeme: `=`, Line: 1},
				{TokenType: token.AT, Lexeme: `@`, Line: 1},
				{TokenType: token.LPAREN, Lexeme: `(`, Line: 1},
				{TokenType: token.IDENT, Lexeme: `x`, Line: 1},
				{TokenType: token.COMMA, Lexeme: `,`, Line: 1},
				{TokenType: token.IDENT, Lexeme: `y`, Line: 1},
				{TokenType: token.COMMA, Lexeme: `,`, Line: 1},
				{TokenType: token.IDENT, Lexeme: `z`, Line: 1},
				{TokenType: token.RPAREN, Lexeme: `)`, Line: 1},
				{TokenType: token.LPAREN, Lexeme: `(`, Line: 1},
				{TokenType: token.IDENT, Lexeme: `x`, Line: 1},
				{TokenType: token.ADD, Lexeme: `+`, Line: 1},
				{TokenType: token.IDENT, Lexeme: `y`, Line: 1},
				{TokenType: token.ADD, Lexeme: `+`, Line: 1},
				{TokenType: token.IDENT, Lexeme: `z`, Line: 1},
				{TokenType: token.RPAREN, Lexeme: `)`, Line: 1},
				{TokenType: token.SEMICOLON, Lexeme: `;`, Line: 1},
				{TokenType: token.EOF, Lexeme: ``, Line: 1},
			},
		},
	}
	for _, tc := range tests {
		s := ScanFile(tc.sourceFile)
		require.Len(t, s.tokens, len(tc.expToks), tc.name)
		for i := range tc.expToks {
			assert.Equal(t, tc.expToks[i], s.tokens[i], tc.name)
		}
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
