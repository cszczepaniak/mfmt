package scanner

import (
	"testing"

	"github.com/cszczepaniak/mfmt/token"
	"github.com/stretchr/testify/assert"
)

func TestScanFile(t *testing.T) {
	tests := []struct {
		name        string
		sourceFile  string
		expTokTypes []token.Type
	}{
		{
			name:        "test simple file",
			sourceFile:  "testdata/simple.m",
			expTokTypes: []token.Type{token.INT, token.MUL, token.INT, token.ADD, token.FLOAT, token.EOF},
		},
	}
	for _, tc := range tests {
		s := ScanFile(tc.sourceFile)
		actTokTypes := make([]token.Type, 0)
		for _, t := range s.tokens {
			actTokTypes = append(actTokTypes, t.TokenType)
		}
		assert.Equal(t, tc.expTokTypes, actTokTypes, tc.name)
	}
}
