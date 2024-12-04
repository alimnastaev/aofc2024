package day4_test

// https://adventofcode.com/2024/day/4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_day4(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		want     int
		testFunc func(string) int
	}{
		{"Day 4: Part 1 - Example", "example1.txt", 0, day4Part1},
		{"Day 4: Part 1 - Input", "input.txt", 0, day4Part1},
		{"Day 4: Part 2 - Example", "example2.txt", 0, day4Part2},
		{"Day 4: Part 2 - Input", "input.txt", 0, day4Part2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

func day4Part1(path string) int {
	return 0
}

func day4Part2(path string) int {
	return 0
}
