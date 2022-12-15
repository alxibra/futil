package futil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSqlInt(t *testing.T) {
	var tests = []struct {
		name  string
		exp   string
		given int
	}{
		{"2 parameter", "(?,?)", 2},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			act := InSqlInt(tt.given)
			assert.Equal(t, tt.exp, act)
		})
	}
}
