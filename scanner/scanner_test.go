package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanner_advance(t *testing.T) {
	tests := []struct {
		name   string
		source string
		s      *Scanner
	}{
		{
			name:   "test next",
			source: "abcdefghi",
			s:      NewScanner("abcdefghi"),
		},
	}
	for _, tt := range tests {
		for _, c := range tt.source {
			assert.Equal(t, c, tt.s.ch)
			tt.s.advance()
		}
	}
}
