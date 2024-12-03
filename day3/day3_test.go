package day3_test

// https://adventofcode.com/2024/day/3

import (
	"regexp"
	"testing"

	"github.com/alimnastaev/aofc2023/utils"
	"github.com/stretchr/testify/require"
)

func Test_day3(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		want     int
		testFunc func(string) int
	}{
		{"Day 3: Part 1 - Example", "example1.txt", 161, day3Part1},
		{"Day 3: Part 1 - Input", "input.txt", 178794710, day3Part1},
		{"Day 3: Part 2 - Example", "example2.txt", 48, day3Part2},
		{"Day 3: Part 2 - Input", "input.txt", 692, day3Part2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

var mulRegex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func day3Part1(path string) int {
	var result int
	for _, line := range utils.ReadFile(path) {
		for _, match := range mulRegex.FindAllString(line, -1) {
			result += extractAndMultiply(match)
		}
	}

	return result
}

func extractAndMultiply(match string) int {
	parts := mulRegex.FindStringSubmatch(match)
	if len(parts) < 3 {
		return 0
	}

	return utils.ParseInt(parts[1]) * utils.ParseInt(parts[2])
}

func day3Part2(path string) int {
	file := utils.ReadFile(path)

	var result int
	for _, line := range file {
		_ = line
	}

	return result
}
