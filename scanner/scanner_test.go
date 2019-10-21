package scanner

import (
	"fmt"
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
		for i, c := range tt.source {
			assert.Equal(t, c, tt.s.advance(), fmt.Sprintf("failed at index %d", i))
		}
	}
}
