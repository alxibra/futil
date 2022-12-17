package futil

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInSQL(t *testing.T) {
	var tests = []struct {
		name  string
		exp   string
		given []string
	}{
		{"2 parameter", `("1","2")`, []string{"1", "2"}},
		{"1 parameter", `("1")`, []string{"1"}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			act := InSQL(tt.given)
			assert.Equal(t, tt.exp, act)
		})
	}

	var tests2 = []struct {
		name  string
		exp   string
		given []int
	}{
		{"2 parameter int", "(1,2)", []int{1, 2}},
		{"1 parameter int", "(1)", []int{1}},
	}
	for _, tt := range tests2 {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			act := InSQL(tt.given)
			assert.Equal(t, tt.exp, act)
		})
	}
}

func TestBenchmark(t *testing.T) {
	stop := Benchmark()
	time.Sleep(100 * time.Millisecond)
	fmt.Println(stop())
}
