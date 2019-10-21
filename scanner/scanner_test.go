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
	}
	for _, tt := range tests {
		s := NewScanner(tt.source)
		s.scanToken()
		tok := s.tokens[0]
		assert.Equal(t, tt.expect, tok)
	}
}
