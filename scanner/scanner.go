package scanner

// Scanner stores state
type Scanner struct {
	source  []rune
	start   int
	current int
}

// Scanner needs to keep track of its current position in the file,
// and should also define functions used to scan various token types in
// a MATLAB source file.

// NewScanner instantiates a scanner
func NewScanner(source string) *Scanner {
	var scanner Scanner
	scanner.source = []rune(source)
	scanner.start = 0
	return &scanner
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
