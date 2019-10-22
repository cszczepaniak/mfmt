package scanner

import (
	"testing"

	"github.com/cszczepaniak/mfmt/token"
	"github.com/stretchr/testify/assert"
)

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

/* These tests will be part of testing scanning multiple tokens that I'll add later
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
	// This will scan to be an INT and an ILLEGAL
	name:   "test illegal number",
	source: "12.",
	expect: token.Token{TokenType: token.INT, Lexeme: "12", Line: 1},
},
{
	// This will scan to be a FLOAT and an IDENT, which the parser will flag later
	name:   "test illegal float",
	source: "12.1234e",
	expect: token.Token{TokenType: token.FLOAT, Lexeme: "12.1234", Line: 1},
},
{
	// This will scan to be an INT and an IDENT, which the parser will flag later
	name:   "test illegal number",
	source: "12e",
	expect: token.Token{TokenType: token.INT, Lexeme: "12", Line: 1},
},
*/

func TestScanner_scanNumber(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		expTokType token.Type
	}{
		{
			name:       "test int",
			source:     "1234",
			expTokType: token.INT,
		},
		{
			name:       "test float",
			source:     "12.34",
			expTokType: token.FLOAT,
		},
		{
			name:       "test no leading zero",
			source:     ".34",
			expTokType: token.FLOAT,
		},
		{
			name:       "test exp",
			source:     "1e4",
			expTokType: token.FLOAT,
		},
		{
			name:       "test decimal exp",
			source:     "1.1E4",
			expTokType: token.FLOAT,
		},
		{
			name:       "test neg exp",
			source:     "1.1e-4",
			expTokType: token.FLOAT,
		},
		{
			name:       "test complex",
			source:     "5j",
			expTokType: token.COMPLEX,
		},
		{
			name:       "test decimal complex",
			source:     "5.1i",
			expTokType: token.COMPLEX,
		},
		{
			name:       "test exp complex",
			source:     "5.1e-2i",
			expTokType: token.COMPLEX,
		},
	}
	for _, tc := range tests {
		s := NewScanner(tc.source)
		tok, err := s.scanNumber()
		assert.Nil(t, err, tc.name)
		assert.Equal(t, tc.expTokType, tok.TokenType, tc.name)
		assert.Equal(t, tc.source, tok.Lexeme, tc.name)
	}
}
func TestErrsScanner_scanNumber(t *testing.T) {
	tests := []struct {
		name   string
		source string
		expErr string
	}{
		{
			name:   "test trailing .",
			source: "1234.",
			expErr: "Invalid number literal",
		},
		{
			name:   "test trailing e",
			source: "1234e",
			expErr: "Invalid number literal",
		},
	}
	for _, tc := range tests {
		s := NewScanner(tc.source)
		_, err := s.scanNumber()
		assert.Error(t, err, tc.name)
	}
}
