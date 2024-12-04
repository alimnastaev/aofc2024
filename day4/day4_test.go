package day4_test

// https://adventofcode.com/2024/day/3

import (
	"regexp"
	"strings"
	"testing"

	"github.com/alimnastaev/aofc2023/utils"
	"github.com/stretchr/testify/require"
)

func Test_day4(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		want     int
		testFunc func(string) int
	}{
		{"Day 3: Part 1 - Example", "example1.txt", 161, day4Part1},
		{"Day 3: Part 1 - Input", "input.txt", 178794710, day4Part1},
		{"Day 3: Part 2 - Example", "example2.txt", 48, day4Part2},
		{"Day 3: Part 2 - Input", "input.txt", 76729637, day4Part2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

var (
	mulRegex         = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	rejectDontsRegex = regexp.MustCompile(`don't\(\).*?do\(\)`)
)

func day4Part1(path string) int {
	return extractAndMultiply(strings.Join(utils.ReadFile(path), ""))
}

func day4Part2(path string) int {
	cleanedLine := rejectDontsRegex.ReplaceAllString(strings.Join(utils.ReadFile(path), ""), "")
	return extractAndMultiply(cleanedLine)
}

func extractAndMultiply(s string) int {
	var result int
	for _, match := range mulRegex.FindAllStringSubmatch(s, -1) {
		result += utils.ParseInt(match[1]) * utils.ParseInt(match[2])
	}
	return result
}
