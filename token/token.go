// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2019 Connor Szczepaniak
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
//
package token

import (
	"unicode"
)

// Type is an enum for the set of tokens
type Type int

// Token represents the set of lexical tokens of the MATLAB programming language.
type Token struct {
	tokenType Type
	lexeme    string
	line      int
}

// The list of tokens.
const (
	// Special tokens
	ILLEGAL Type = iota
	EOF
	COMMENT

	literal_beg
	// Identifiers and basic type literals
	// (these tokens stand for classes of literals)
	IDENT   // main
	INT     // 12345
	FLOAT   // 123.45
	COMPLEX // 123.45i
	CHAR    // 'abc'
	STRING  // "abc"
	literal_end

	operator_beg
	// Operators and delimiters
	// source: https://www.mathworks.com/help/matlab/matlab_prog/matlab-operators-and-special-characters.html
	ADD         // +
	SUB         // -
	ELEM_MUL    // .*
	MUL         // *
	ELEM_RDIV   // ./
	RDIV        // /
	ELEM_LDIV   // .\
	LDIV        // \
	ELEM_PWR    // .^
	PWR         // ^
	TRANSP      // .'
	COMP_TRANSP // '

	EQL // ==
	NEQ // ~=
	GTR // >
	GEQ // >=
	LSS // <
	LEQ // <=

	AND  // &
	OR   // |
	LAND // &&
	LOR  // ||
	NOT  // ~

	AT // @

	ASSIGN // =
	ARROW  // <-

	ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	operator_end

	keyword_beg
	// Keywords
	BREAK
	CASE
	CATCH
	CLASSDEF
	CONTINUE
	ELSE
	ELSEIF
	END
	FOR
	FUNCTION
	GLOBAL
	IF
	OTHERWISE
	PARFOR
	PERSISTENT
	RETURN
	SPMD
	SWITCH
	TRY
	WHILE
	keyword_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",

	EOF:     "EOF",
	COMMENT: "COMMENT",

	IDENT:   "IDENT",
	INT:     "INT",
	FLOAT:   "FLOAT",
	COMPLEX: "COMPLEX",
	CHAR:    "CHAR",
	STRING:  "STRING",

	ADD:         "+",
	SUB:         "-",
	ELEM_MUL:    ".*",
	MUL:         "*",
	ELEM_RDIV:   "./",
	RDIV:        "/",
	ELEM_LDIV:   ".\\",
	LDIV:        "\\",
	ELEM_PWR:    ".^",
	PWR:         "^",
	TRANSP:      ".'",
	COMP_TRANSP: "'",

	EQL: "==",
	NEQ: "~=",
	GTR: ">",
	GEQ: ">=",
	LSS: "<",
	LEQ: "<=",

	AND:  "&",
	OR:   "|",
	LAND: "&&",
	LOR:  "||",
	NOT:  "~",

	AT: "@",

	ASSIGN: "=",

	ELLIPSIS: "...",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",

	BREAK:      "break",
	CASE:       "case",
	CATCH:      "catch",
	CLASSDEF:   "classdef",
	CONTINUE:   "continue",
	ELSE:       "else",
	ELSEIF:     "elseif",
	END:        "end",
	FOR:        "for",
	FUNCTION:   "function",
	GLOBAL:     "global",
	IF:         "if",
	OTHERWISE:  "otherwise",
	PARFOR:     "parfor",
	PERSISTENT: "persistent",
	RETURN:     "return",
	SPMD:       "spmd",
	SWITCH:     "switch",
	TRY:        "try",
	WHILE:      "while",
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token ADD, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token IDENT, the string is "IDENT").
//
func (tok Token) String() string {
	return tok.lexeme
}

/* Since mfmt is just a formatter, we don't care about precedence...
// A set of constants for precedence-based expression parsing.
// Non-operators have lowest precedence, followed by operators
// starting with precedence 1 up to unary operators. The highest
// precedence serves as "catch-all" precedence for selector,
// indexing, and other operator and delimiter tokens.
//
const (
	LowestPrec  = 0 // non-operators
	UnaryPrec   = 6
	HighestPrec = 7
)

// Precedence returns the operator precedence of the binary
// operator op. If op is not a binary operator, the result
// is LowestPrecedence.
//
func (op Token) Precedence() int {
	switch op {
	case LOR:
		return 1
	case LAND:
		return 2
	case EQL, NEQ, LSS, LEQ, GTR, GEQ:
		return 3
	case ADD, SUB, OR, XOR:
		return 4
	case MUL, QUO, REM, SHL, SHR, AND, AND_NOT:
		return 5
	}
	return LowestPrec
}

*/

var keywords map[string]Type

func init() {
	keywords = make(map[string]Type)
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
}

// Lookup maps an identifier to its keyword token or IDENT (if not a keyword).
//
func Lookup(ident string) Type {
	if tokType, isKeyword := keywords[ident]; isKeyword {
		return tokType
	}
	return IDENT
}

// Predicates

// IsLiteral returns true for tokens corresponding to identifiers
// and basic type literals; it returns false otherwise.
//
func (tok Token) IsLiteral() bool { return literal_beg < tok.tokenType && tok.tokenType < literal_end }

// IsOperator returns true for tokens corresponding to operators and
// delimiters; it returns false otherwise.
//
func (tok Token) IsOperator() bool {
	return operator_beg < tok.tokenType && tok.tokenType < operator_end
}

// IsKeyword returns true for tokens corresponding to keywords;
// it returns false otherwise.
//
func (tok Token) IsKeyword() bool { return keyword_beg < tok.tokenType && tok.tokenType < keyword_end }

// IsKeyword reports whether name is a Go keyword, such as "function" or "return".
//
func IsKeyword(name string) bool {
	// TODO: opt: use a perfect hash function instead of a global map.
	_, ok := keywords[name]
	return ok
}

// IsIdentifier reports whether name is a MATLAB identifier, that is, a non-empty
// string made up of letters, digits, and underscores, where the first character
// is not a digit or underscore. Keywords are not identifiers.
//
func IsIdentifier(name string) bool {
	for i, c := range name {
		if !unicode.IsLetter(c) && (i == 0 || !unicode.IsDigit(c) && c != '_') {
			return false
		}
	}
	return name != "" && !IsKeyword(name)
}
