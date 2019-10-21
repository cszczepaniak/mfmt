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

// advance consumes the current character
func (s *Scanner) advance() rune {
	ch := s.source[s.current]
	if s.current < len(s.source)-1 {
		s.current++
	}
	return ch
}
