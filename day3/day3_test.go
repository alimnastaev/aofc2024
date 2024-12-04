package day3_test

// https://adventofcode.com/2024/day/3

import (
	"regexp"
	"strings"
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
		{"Day 3: Part 2 - Input", "input.txt", 76729637, day3Part2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

func day3Part1(path string) int {
	return extractAndMultiply(strings.Join(utils.ReadFile(path), ""))
}

func day3Part2(path string) int {
	cleanedLine := regexp.
		MustCompile(`don't\(\).*?do\(\)`).
		ReplaceAllString(strings.Join(utils.ReadFile(path), ""), "")
	return extractAndMultiply(cleanedLine)
}

func extractAndMultiply(s string) int {
	var result int
	for _, match := range regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`).FindAllStringSubmatch(s, -1) {
		result += utils.ParseInt(match[1]) * utils.ParseInt(match[2])
	}
	return result
}
