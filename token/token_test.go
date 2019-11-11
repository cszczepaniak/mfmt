// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2019 Connor Szczepaniak
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
//

package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookup(t *testing.T) {
	tests := []struct {
		name  string
		ident string
		want  Type
	}{
		{
			name:  "test keyword",
			ident: "function",
			want:  FUNCTION,
		},
		{
			name:  "test identifier",
			ident: "something",
			want:  IDENT,
		},
	}
	for _, tt := range tests {
		got := LookupIdent(tt.ident)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestToken_IsLiteral(t *testing.T) {
	tests := []struct {
		name string
		tok  Token
		want bool
	}{
		{
			name: "test literal",
			tok:  Token{TokenType: CHAR, Lexeme: "abcd", Line: 0},
			want: true,
		},
		{
			name: "test non-literal",
			tok:  Token{TokenType: ADD, Lexeme: "+", Line: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		got := tt.tok.IsLiteral()
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestToken_IsOperator(t *testing.T) {
	tests := []struct {
		name string
		tok  Token
		want bool
	}{
		{
			name: "test operator",
			tok:  Token{TokenType: LDIV, Lexeme: "\\", Line: 0},
			want: true,
		},
		{
			name: "test non-operator",
			tok:  Token{TokenType: EOF, Lexeme: "", Line: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		got := tt.tok.IsOperator()
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestToken_IsKeyword(t *testing.T) {
	tests := []struct {
		name string
		tok  Token
		want bool
	}{
		{
			name: "test keyword",
			tok:  Token{TokenType: IF, Lexeme: "if", Line: 0},
			want: true,
		},
		{
			name: "test non-keyword",
			tok:  Token{TokenType: ADD, Lexeme: "+", Line: 0},
			want: false,
		},
	}
	for _, tt := range tests {
		got := tt.tok.IsKeyword()
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestIsKeyword(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "test keyword",
			str:  "for",
			want: true,
		},
		{
			name: "test non-keyword",
			str:  "something",
			want: false,
		},
	}
	for _, tt := range tests {
		got := IsKeyword(tt.str)
		assert.Equal(t, tt.want, got, tt.name)
	}
}

func TestIsIdentifier(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "test alphanumeric",
			str:  "azAZ09",
			want: true,
		},
		{
			name: "test alphanumeric with underscore",
			str:  "az_AZ_09",
			want: true,
		},
		{
			name: "test leading number",
			str:  "09azAZ",
			want: false,
		},
		{
			name: "test leading underscore",
			str:  "_azAZ09",
			want: false,
		},
		{
			name: "test special character",
			str:  "azAZ09!",
			want: false,
		},
	}
	for _, tt := range tests {
		got := IsIdentifier(tt.str)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
