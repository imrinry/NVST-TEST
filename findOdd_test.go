package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findOddInt(t *testing.T) {

	type testCase struct {
		name     string
		input    []int
		expected int
	}

	cases := []testCase{
		{name: "7", input: []int{7}, expected: 7},
		{name: "0", input: []int{0}, expected: 0},
		{name: "1,1,2", input: []int{1, 1, 2}, expected: 2},
		{name: "0,1,0,1,0", input: []int{0, 1, 0, 1, 0}, expected: 0},
		{name: "1,2,2,3,3,3,4,3,3,3,2,2,1", input: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, expected: 4},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := findOddInt(c.input)
			assert.Equal(t, c.expected, input)
		})
	}
}

func Benchmark_findOddInt(b *testing.B) {

	for i := 0; i < b.N; i++ {
		findOddInt([]int{7})
	}

}
