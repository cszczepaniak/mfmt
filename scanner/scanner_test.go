package scanner

import (
	"fmt"
	"testing"

	"github.com/cszczepaniak/mfmt/token"
	"github.com/stretchr/testify/assert"
)

func TestScanner_advance(t *testing.T) {
	tests := []struct {
		name   string
		source string
		s      *Scanner
	}{
		{
			name:   "test next",
			source: "abcdefghi",
			s:      NewScanner("abcdefghi"),
		},
	}
	for _, tt := range tests {
		for i, c := range tt.source {
			assert.Equal(t, c, tt.s.advance(), fmt.Sprintf("failed at index %d", i))
		}
	}
}

func TestScanner_peek(t *testing.T) {
	tests := []struct {
		name   string
		source string
		s      *Scanner
	}{
		{
			name:   "test peek",
			source: "abcdefghi",
			s:      NewScanner("abcdefghi"),
		},
	}
	for _, tt := range tests {
		for i := range tt.source {
			assert.Equal(t, 'a', tt.s.peek(), fmt.Sprintf("failed at index %d", i))
		}
	}
}

func Test_isDigit(t *testing.T) {
	tests := []struct {
		name   string
		in     rune
		expect bool
	}{
		{
			name:   "test digit",
			in:     '5',
			expect: true,
		},
		{
			name:   "test nondigit",
			in:     'a',
			expect: false,
		},
	}
	for _, tt := range tests {
		got := isDigit(tt.in)
		assert.Equal(t, tt.expect, got)
	}
}

func Test_isAlpha(t *testing.T) {
	tests := []struct {
		name   string
		in     rune
		expect bool
	}{
		{
			name:   "test alpha",
			in:     'a',
			expect: true,
		},
		{
			name:   "test alpha",
			in:     'Z',
			expect: true,
		},
		{
			name:   "test nonalpha",
			in:     '_',
			expect: false,
		},
	}
	for _, tt := range tests {
		got := isAlpha(tt.in)
		assert.Equal(t, tt.expect, got)
	}
}

func TestScanner_scanToken(t *testing.T) {
	tests := []struct {
		name   string
		source string
		expect token.Token
	}{
		{
			name:   "test +",
			source: "+",
			expect: token.Token{TokenType: token.ADD, Lexeme: "+", Line: 1},
		},
		{
			name:   "test -",
			source: "-",
			expect: token.Token{TokenType: token.SUB, Lexeme: "-", Line: 1},
		},
		{
			name:   "test *",
			source: "*",
			expect: token.Token{TokenType: token.MUL, Lexeme: "*", Line: 1},
		},
		{
			name:   "test <",
			source: "<",
			expect: token.Token{TokenType: token.LSS, Lexeme: "<", Line: 1},
		},
		{
			name:   "test <=",
			source: "<=",
			expect: token.Token{TokenType: token.LEQ, Lexeme: "<=", Line: 1},
		},
		{
			name:   "test .*",
			source: ".*",
			expect: token.Token{TokenType: token.ELEM_MUL, Lexeme: ".*", Line: 1},
		},
		{
			name:   "test ...",
			source: "...",
			expect: token.Token{TokenType: token.ELLIPSIS, Lexeme: "...", Line: 1},
		},
		{
			name:   "test ..",
			source: "..",
			expect: token.Token{TokenType: token.ILLEGAL, Lexeme: "..", Line: 1},
		},
		{
			name:   "test whitespace",
			source: "\t\r\v+",
			expect: token.Token{TokenType: token.ADD, Lexeme: "+", Line: 1},
		},
		{
			name:   "test identifier",
			source: "myVar",
			expect: token.Token{TokenType: token.IDENT, Lexeme: "myVar", Line: 1},
		},
		{
			name:   "test keyword",
			source: "function",
			expect: token.Token{TokenType: token.FUNCTION, Lexeme: "function", Line: 1},
		},
		{
			// This will scan to be an INT and an IDENT, but the parser will flag it later
			name:   "test illegal identifier",
			source: "3real5me",
			expect: token.Token{TokenType: token.INT, Lexeme: "3", Line: 1},
		},
		{
			// This will scan to be an ILLEGAL and then an IDENT
			name:   "test illegal identifier",
			source: "_myVar",
			expect: token.Token{TokenType: token.ILLEGAL, Lexeme: "_", Line: 1},
		},
		{
			name:   "test illegal number",
			source: "12.",
			expect: token.Token{TokenType: token.ILLEGAL, Lexeme: "12.", Line: 1},
		},
		{
			name:   "test illegal number",
			source: "12e",
			expect: token.Token{TokenType: token.ILLEGAL, Lexeme: "12e", Line: 1},
		},
		{
			name:   "test float",
			source: "12.1234",
			expect: token.Token{TokenType: token.FLOAT, Lexeme: "12.1234", Line: 1},
		},
		{
			name:   "test no leading zero",
			source: ".1234",
			expect: token.Token{TokenType: token.FLOAT, Lexeme: ".1234", Line: 1},
		},
		{
			name:   "test illegal float",
			source: "12.1234e",
			expect: token.Token{TokenType: token.ILLEGAL, Lexeme: "12.1234e", Line: 1},
		},
		{
			name:   "test int",
			source: "1234",
			expect: token.Token{TokenType: token.INT, Lexeme: "1234", Line: 1},
		},
		{
			name:   "test scientific",
			source: "1.3e4",
			expect: token.Token{TokenType: token.FLOAT, Lexeme: "1.3e4", Line: 1},
		},
		{
			name:   "test scientific",
			source: "1e4",
			expect: token.Token{TokenType: token.FLOAT, Lexeme: "1e4", Line: 1},
		},
		{
			name:   "test complex",
			source: "12i",
			expect: token.Token{TokenType: token.COMPLEX, Lexeme: "12i", Line: 1},
		},
		{
			name:   "test scientific complex",
			source: "1.3e4j",
			expect: token.Token{TokenType: token.COMPLEX, Lexeme: "1.3e4j", Line: 1},
		},
	}
	for _, tt := range tests {
		s := NewScanner(tt.source)
		for !s.isAtEnd() {
			s.scanToken()
		}
		tok := s.tokens[0]
		assert.Equal(t, tt.expect, tok, tt.name)
	}
}
