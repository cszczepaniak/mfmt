package scanner

import "github.com/cszczepaniak/mfmt/token"

// Scanner stores state
type Scanner struct {
	source  []rune
	start   int
	current int
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
	scanner.start, scanner.current, scanner.line = 0, 0, 1
	scanner.tokens = make([]token.Token, 0)
	return &scanner
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) scanToken() {
	s.start = s.current
	c := s.advance()
	switch c {
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
		case isAlpha(c):
			s.scanWord()
		case isDigit(c):
			s.scanNumber()
		default:
			s.tokens = append(s.tokens, s.makeToken(token.ILLEGAL))
		}
	}
}

func (s *Scanner) scanWord() {
	c := s.peek()
	for !s.isAtEnd() && (isAlpha(c) || isDigit(c) || c == '_') {
		s.advance()
	}
	word := string(s.source[s.start:s.current])
	tokType := token.Lookup(word)
	s.tokens = append(s.tokens, s.makeToken(tokType))
}

// TODO this might just be the ugliest function I've ever written..... but it passes the tests.
// clean it up later............
func (s *Scanner) scanNumber() {
	// Integer part
	for !s.isAtEnd() && isDigit(s.peek()) {
		s.advance()
	}
	// Check for fractional part, scientific notation, imaginary numbers
	switch s.peek() {
	case '.':
		if isDigit(s.peekNext()) {
			// Consume the dot
			s.advance()
			// Consume the next digits
			for !s.isAtEnd() && isDigit(s.peek()) {
				s.advance()
			}
			// Check for any scientific notation
			if s.peek() == 'e' || s.peek() == 'E' {
				if isDigit(s.peekNext()) {
					s.advance()
					for !s.isAtEnd() && isDigit(s.peek()) {
						s.advance()
					}
					s.tokens = append(s.tokens, s.makeToken(token.FLOAT))
				} else {
					s.advance()
					s.tokens = append(s.tokens, s.makeToken(token.ILLEGAL))
				}
			} else {
				s.tokens = append(s.tokens, s.makeToken(token.FLOAT))
			}
		} else {
			// a trailing . is illegal
			s.advance()
			s.tokens = append(s.tokens, s.makeToken(token.ILLEGAL))
		}
	case 'e', 'E':
		if isDigit(s.peekNext()) {
			s.advance()
			for !s.isAtEnd() && isDigit(s.peek()) {
				s.advance()
			}
			s.tokens = append(s.tokens, s.makeToken(token.FLOAT))
		} else {
			// a trailing e or E is illegal
			s.advance()
			s.tokens = append(s.tokens, s.makeToken(token.ILLEGAL))
		}
	case 'i', 'j':
		s.advance()
		s.tokens = append(s.tokens, s.makeToken(token.COMPLEX))
	default:
		s.tokens = append(s.tokens, s.makeToken(token.INT))
	}
}

func (s *Scanner) scanDot() {
	switch {
	case s.match('*'):
		s.tokens = append(s.tokens, s.makeToken(token.ELEM_MUL))
	case s.match('/'):
		s.tokens = append(s.tokens, s.makeToken(token.ELEM_RDIV))
	case s.match('\\'):
		s.tokens = append(s.tokens, s.makeToken(token.ELEM_LDIV))
	case s.match('^'):
		s.tokens = append(s.tokens, s.makeToken(token.ELEM_PWR))
	case s.match('\''):
		s.tokens = append(s.tokens, s.makeToken(token.TRANSP))
	default:
		// Check for ellipsis
		if s.match('.') && s.match('.') {
			s.tokens = append(s.tokens, s.makeToken(token.ELLIPSIS))
		} else {
			s.tokens = append(s.tokens, s.makeToken(token.ILLEGAL))
		}
	}
}

func (s *Scanner) makeToken(tokenType token.Type) token.Token {
	str := s.source[s.start:s.current]
	return token.Token{
		TokenType: tokenType,
		Lexeme:    string(str),
		Line:      s.line,
	}
}

// isAtEnd checks if current is pointing at the end of the file
func (s *Scanner) isAtEnd() bool {
	return s.current == len(s.source)
}

// peek looks at the current character without consuming it
func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

// advance consumes and returns the current character
func (s *Scanner) advance() rune {
	if !s.isAtEnd() {
		s.current++
	}
	return s.source[s.current-1]
}

// match is like advance, but the current character must match a condition first
func (s *Scanner) match(c rune) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != c {
		return false
	}
	s.current++
	return true
}
