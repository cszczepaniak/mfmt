package scanner

import (
	"testing"

	"github.com/cszczepaniak/mfmt/token"
	"github.com/stretchr/testify/assert"
)

func TestScanner_scanNumber(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		expTokType token.Type
		expLexeme  string
	}{
		{
			name:       "test int with trailing nonint char",
			source:     "1234+",
			expTokType: token.INT,
			expLexeme:  "1234",
		},
		{
			name:       "test int",
			source:     "1234",
			expTokType: token.INT,
			expLexeme:  "1234",
		},
		{
			name:       "test float",
			source:     "12.34",
			expTokType: token.FLOAT,
			expLexeme:  "12.34",
		},
		{
			name:       "test no leading zero",
			source:     ".34",
			expTokType: token.FLOAT,
			expLexeme:  ".34",
		},
		{
			name:       "test exp",
			source:     "1e4",
			expTokType: token.FLOAT,
			expLexeme:  "1e4",
		},
		{
			name:       "test decimal exp",
			source:     "1.1E4",
			expTokType: token.FLOAT,
			expLexeme:  "1.1E4",
		},
		{
			name:       "test neg exp",
			source:     "1.1e-4",
			expTokType: token.FLOAT,
			expLexeme:  "1.1e-4",
		},
		{
			name:       "test complex",
			source:     "5j",
			expTokType: token.COMPLEX,
			expLexeme:  "5j",
		},
		{
			name:       "test decimal complex",
			source:     "5.1i",
			expTokType: token.COMPLEX,
			expLexeme:  "5.1i",
		},
		{
			name:       "test exp complex",
			source:     "5.1e-2i",
			expTokType: token.COMPLEX,
			expLexeme:  "5.1e-2i",
		},
	}
	for _, tc := range tests {
		s := New(tc.source)
		tok, err := s.scanNumber()
		assert.Nil(t, err, tc.name)
		assert.Equal(t, tc.expTokType, tok.TokenType, tc.name)
		assert.Equal(t, tc.expLexeme, tok.Lexeme, tc.name)
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
		s := New(tc.source)
		_, err := s.scanNumber()
		assert.Error(t, err, tc.name)
	}
}

func TestScanner_scanWord(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		expTokType token.Type
		expLexeme  string
	}{
		{
			name:       "test identifier followed by something else",
			source:     "abc123 = 5;",
			expTokType: token.IDENT,
			expLexeme:  "abc123",
		},
		{
			name:       "test keyword followed by something else",
			source:     "if myBool == 3",
			expTokType: token.IF,
			expLexeme:  "if",
		},
		{
			name:       "test identifier",
			source:     "abc123",
			expTokType: token.IDENT,
			expLexeme:  "abc123",
		},
		{
			name:       "test identifier",
			source:     "abc_123",
			expTokType: token.IDENT,
			expLexeme:  "abc_123",
		},
		{
			name:       "test keyword",
			source:     "classdef",
			expTokType: token.CLASSDEF,
			expLexeme:  "classdef",
		},
	}
	for _, tc := range tests {
		s := New(tc.source)
		tok, err := s.scanWord()
		assert.Nil(t, err, tc.name)
		assert.Equal(t, tc.expTokType, tok.TokenType, tc.name)
		assert.Equal(t, tc.expLexeme, tok.Lexeme, tc.name)
	}
}

func TestErrsScanner_scanWord(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		expTokType token.Type
	}{
		{
			name:       "test invalid identifier",
			source:     "_abc123",
			expTokType: token.IDENT,
		},
	}
	for _, tc := range tests {
		s := New(tc.source)
		_, err := s.scanWord()
		assert.Error(t, err, tc.name)
	}
}

func TestScanner_scanDot(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		expTokType token.Type
		expError   bool
		expLexeme  string
	}{
		{
			name:       "test element wise multiplication",
			source:     ".*",
			expTokType: token.ELEM_MUL,
			expError:   false,
			expLexeme:  ".*",
		},
		{
			name:       "test ellipsis",
			source:     "...",
			expTokType: token.ELLIPSIS,
			expError:   false,
			expLexeme:  "...",
		},
		{
			name:       "test two dots",
			source:     "..",
			expTokType: 0,
			expError:   true,
			expLexeme:  "",
		},
		{
			name:       "test single dot",
			source:     ".a",
			expTokType: token.PERIOD,
			expError:   false,
			expLexeme:  ".",
		},
	}
	for _, tc := range tests {
		s := New(tc.source)
		tok, err := s.scanDot()
		if tc.expError {
			assert.Error(t, err, tc.name)
		} else {
			assert.Nil(t, err, tc.name)
			assert.Equal(t, tc.expTokType, tok.TokenType, tc.name)
			assert.Equal(t, tc.expLexeme, tok.Lexeme, tc.name)
		}
	}
}

func TestScanner_scanString(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		expTokType token.Type
	}{
		{
			name:       "test string",
			source:     "\"abc123\"",
			expTokType: token.STRING,
		},
		{
			name:       "test another string",
			source:     "\"abc123\"",
			expTokType: token.STRING,
		},
	}
	for _, tc := range tests {
		s := New(tc.source)
		tok, err := s.scanString()
		assert.Nil(t, err, tc.name)
		assert.Equal(t, tc.expTokType, tok.TokenType, tc.name)
		assert.Equal(t, tc.source, tok.Lexeme, tc.name)
	}
}

func TestErrsScanner_scanString(t *testing.T) {
	tests := []struct {
		name   string
		source string
		expErr string
	}{
		{
			name:   "test unterminated string",
			source: "\"abc",
			expErr: "Unterminated string literal",
		},
		{
			name:   "test new line in string",
			source: "\"a\na\"",
			expErr: "Unterminated string literal",
		},
	}
	for _, tc := range tests {
		s := New(tc.source)
		_, err := s.scanString()
		assert.Error(t, err, tc.name)
	}
}

func TestScanner_consumeDigits(t *testing.T) {
	tests := []struct {
		name         string
		source       string
		expNumDigits int
		expReadIdx   int
	}{
		{
			name:         "test just digits",
			source:       "1234",
			expNumDigits: 4,
			expReadIdx:   4,
		},
		{
			name:         "test digits followed by other stuff",
			source:       "1234.1234",
			expNumDigits: 4,
			expReadIdx:   5,
		},
		{
			name:         "no digits",
			source:       ".1234",
			expNumDigits: 0,
			expReadIdx:   1,
		},
	}
	for _, tc := range tests {
		s := New(tc.source)
		act := s.consumeDigits()
		assert.Equal(t, tc.expNumDigits, act, tc.name, "(num digits)")
		assert.Equal(t, tc.expReadIdx, s.readIdx, tc.name, "(read idx)")
	}
}
