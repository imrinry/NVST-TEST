package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSmileys(t *testing.T) {

	type testCase struct {
		name     string
		input    []string
		expected int
	}

	cases := []testCase{
		{name: "1", input: []string{":)", ";(", ";}", ":-D"}, expected: 2},
		{name: "2", input: []string{";D", ":-(", ":-)", ";~)"}, expected: 3},
		{name: "3", input: []string{";]", ":[", ";*", ":$", ";-D"}, expected: 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := countSmileys(c.input)
			assert.Equal(t, c.expected, input)
		})
	}
}

func Benchmark_countSmileys(b *testing.B) {

	for i := 0; i < b.N; i++ {
		countSmileys([]string{":)", ";(", ";}", ":-D"})
	}

}
