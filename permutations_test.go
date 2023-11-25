package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generatePermutations(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		expected []string
	}

	cases := []testCase{
		{name: "a", input: "a", expected: []string{"a"}},
		{name: "ab", input: "ab", expected: []string{"ab", "ba"}},
		{name: "abc", input: "abc", expected: []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{name: "aabb", input: "aabb", expected: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := generatePermutations(c.input)
			assert.Equal(t, c.expected, input)
		})
	}
}

func Benchmark_generatePermutations(b *testing.B) {

	for i := 0; i < b.N; i++ {
		generatePermutations("aabb")
	}

}
