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
	switch s.advance() {
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
		s.scanDot()
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
	return s.source[s.current]
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
