package scanner

import (
	"errors"
	"unicode"

	"github.com/cszczepaniak/mfmt/token"
)

// Scanner stores state
type Scanner struct {
	source  []rune
	c       rune
	idx     int
	readIdx int
	line    int
	tokens  []token.Token
}

// Scanner needs to keep track of its current position in the file,
// and should also define functions used to scan various token types in
// a MATLAB source file.

// NewScanner instantiates a scanner
func NewScanner(source string) *Scanner {
	var scanner Scanner
	scanner.source = []rune(source)
	scanner.idx, scanner.readIdx, scanner.line = 0, 1, 1
	scanner.tokens = make([]token.Token, 0)
	scanner.c = scanner.source[0]
	return &scanner
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) scanToken() {
	s.idx = s.readIdx
	s.advance()
	switch s.c {
	// One-character tokens are up first
	case '+':
		s.tokens = append(s.tokens, s.makeToken(token.ADD))
	case '-':
		s.tokens = append(s.tokens, s.makeToken(token.SUB))
	case '*':
		s.tokens = append(s.tokens, s.makeToken(token.MUL))
	case '\\':
		s.tokens = append(s.tokens, s.makeToken(token.LDIV))
	case '(':
		s.tokens = append(s.tokens, s.makeToken(token.LPAREN))
	case ')':
		s.tokens = append(s.tokens, s.makeToken(token.RPAREN))
	case '[':
		s.tokens = append(s.tokens, s.makeToken(token.LBRACK))
	case ']':
		s.tokens = append(s.tokens, s.makeToken(token.RBRACK))
	case '{':
		s.tokens = append(s.tokens, s.makeToken(token.LBRACE))
	case ';':
		s.tokens = append(s.tokens, s.makeToken(token.SEMICOLON))
	case ':':
		s.tokens = append(s.tokens, s.makeToken(token.COLON))
	case ',':
		s.tokens = append(s.tokens, s.makeToken(token.COMMA))
	case '\'':
		s.tokens = append(s.tokens, s.makeToken(token.COMMA))
	// Next handle two-character operators
	case '~':
		if s.match('=') {
			s.tokens = append(s.tokens, s.makeToken(token.NEQ))
		} else {
			s.tokens = append(s.tokens, s.makeToken(token.NOT))
		}
	case '=':
		if s.match('=') {
			s.tokens = append(s.tokens, s.makeToken(token.EQL))
		} else {
			s.tokens = append(s.tokens, s.makeToken(token.ASSIGN))
		}
	case '<':
		if s.match('=') {
			s.tokens = append(s.tokens, s.makeToken(token.LEQ))
		} else {
			s.tokens = append(s.tokens, s.makeToken(token.LSS))
		}
	case '>':
		if s.match('=') {
			s.tokens = append(s.tokens, s.makeToken(token.GEQ))
		} else {
			s.tokens = append(s.tokens, s.makeToken(token.GTR))
		}
	case '.':
		// The dot is interesting because of ellipses and number literals with no leading zero
		if isDigit(s.peek()) {
			s.retreat()
			s.scanNumber()
		} else {
			s.scanDot()
		}
	case '\r', '\f', '\t', '\v':
		// Skip whitespace
	case '\n':
		// Handle new lines
		s.line++
	default:
		// Check for literals in default case
		switch {
		case isAlpha(s.c):
			s.scanWord()
		case isDigit(s.c):
			s.scanNumber()
		default:
			s.tokens = append(s.tokens, s.makeToken(token.ILLEGAL))
		}
	}
}

func (s *Scanner) scanWord() (token.Token, error) {
	tokType := token.IDENT
	for !s.isAtEnd() && (isAlpha(s.c) || isDigit(s.c) || s.c == '_') {
		s.advance()
	}
	word := string(s.source[s.idx:s.readIdx])
	if !token.IsIdentifier(word) {
		if token.IsKeyword(word) {
			tokType = token.Lookup(word)
		} else {
			return token.Token{}, errors.New("Invalid identifier")
		}
	}
	return s.makeToken(tokType), nil
}

func (s *Scanner) scanNumber() (token.Token, error) {
	tokType := token.INT
	// Integer part
	s.consumeDigits()
	// Check for fractional part
	if s.c == '.' {
		tokType = token.FLOAT
		s.advance()
		n := s.consumeDigits()
		if n == 0 {
			return token.Token{}, errors.New("Illegal number literal")
		}
	}
	// Check for scientific notation
	if unicode.ToLower(s.c) == 'e' {
		tokType = token.FLOAT
		s.advance()
		if s.c == '-' {
			s.advance()
		}
		n := s.consumeDigits()
		if n == 0 {
			return token.Token{}, errors.New("Illegal number literal")
		}
	}
	// Check for complex
	if s.c == 'i' || s.c == 'j' {
		tokType = token.COMPLEX
		s.advance()
	}
	return s.makeToken(tokType), nil
}

func (s *Scanner) consumeDigits() int {
	i := 0
	for isDigit(s.c) {
		s.advance()
		i++
		if s.isAtEnd() {
			break
		}
	}
	return i
}

func (s *Scanner) scanDot() (token.Token, error) {
	tokType := token.PERIOD
	s.advance()
	switch s.c {
	case '*':
		tokType = token.ELEM_MUL
		s.advance()
	case '/':
		tokType = token.ELEM_RDIV
		s.advance()
	case '\\':
		tokType = token.ELEM_LDIV
		s.advance()
	case '^':
		tokType = token.ELEM_PWR
		s.advance()
	case '\'':
		tokType = token.TRANSP
		s.advance()
	case '.':
		if s.peek() == '.' {
			tokType = token.ELLIPSIS
			s.advance()
			s.advance()
		} else {
			return token.Token{}, errors.New("Syntax error")
		}
	}
	return s.makeToken(tokType), nil
}

func (s *Scanner) scanString() (token.Token, error) {
	// Assume the current character is the opening "
	s.advance()
	for s.peek() != '"' {
		s.advance()
		if s.c == '\n' || s.isAtEnd() {
			return token.Token{}, errors.New("Unterminated string literal")
		}
	}
	s.advance()
	return s.makeToken(token.STRING), nil
}

func (s *Scanner) makeToken(tokenType token.Type) token.Token {
	str := s.source[s.idx:s.readIdx]
	return token.Token{
		TokenType: tokenType,
		Lexeme:    string(str),
		Line:      s.line,
	}
}

// isAtEnd checks if current is pointing at the end of the file
func (s *Scanner) isAtEnd() bool {
	return s.readIdx == len(s.source)
}

// peek looks at the next character without advancing
func (s *Scanner) peek() rune {
	if s.readIdx >= len(s.source) {
		return 0
	}
	return s.source[s.readIdx]
}

// advance consumes and returns the current character
func (s *Scanner) advance() {
	if !s.isAtEnd() {
		c := s.source[s.readIdx]
		s.readIdx++
		s.c = c
	} else {
		s.c = 0
	}
}

func (s *Scanner) retreat() {
	if s.readIdx > 0 {
		s.readIdx--
	}
}

// match is like advance, but the current character must match a condition first
func (s *Scanner) match(c rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.readIdx] != c {
		return false
	}
	s.readIdx++
	return true
}
