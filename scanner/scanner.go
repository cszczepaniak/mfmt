package scanner

import "github.com/cszczepaniak/mfmt/token"

// Scanner stores state
type Scanner struct {
	source  []rune
	start   int
	current int
	tokens  []token.Token
}

// Scanner needs to keep track of its current position in the file,
// and should also define functions used to scan various token types in
// a MATLAB source file.

// NewScanner instantiates a scanner
func NewScanner(source string) *Scanner {
	var scanner Scanner
	scanner.source = []rune(source)
	scanner.start, scanner.current = 0, 0
	scanner.tokens = make([]token.Token, 0)
	return &scanner
}

func isAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
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
