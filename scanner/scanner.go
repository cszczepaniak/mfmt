package scanner

// Scanner stores state
type Scanner struct {
	source  []rune
	ch      rune
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
	scanner.ch = scanner.source[0]
	return &scanner
}

// advance advances the file position
func (s *Scanner) advance() {
	if s.current < len(s.source)-1 {
		s.current++
		s.ch = s.source[s.current]
	}
}

func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\n' || s.ch == '\t' || s.ch == '\r' {
		s.advance()
	}
}
