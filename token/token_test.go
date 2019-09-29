// Copyright 2009 The Go Authors. All rights reserved.
// Copyright 2019 Connor Szczepaniak
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Package token defines constants representing the lexical tokens of the Go
// programming language and basic operations on tokens (printing, predicates).
//

package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken_String(t *testing.T) {
	tests := []struct {
		name string
		tok  Token
		want string
	}{
		{
			name: "operator",
			tok:  ADD,
			want: "+",
		},
		{
			name: "keyword",
			tok:  FUNCTION,
			want: "function",
		},
	}
	for _, tt := range tests {
		got := tt.tok.String()
		assert.Equal(t, tt.want, got)
	}
}
