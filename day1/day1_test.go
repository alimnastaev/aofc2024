package day1_test

// https://adventofcode.com/2024/day/1

import (
	"slices"
	"strings"
	"testing"

	"github.com/alimnastaev/aofc2023/utils"
	"github.com/stretchr/testify/require"
)

func Test_day1(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		want     int
		testFunc func(string) int
	}{
		{"Day 1: Part 1 - Example", "example.txt", 11, day1Part1},
		{"Day 1: Part 1 - Input", "input.txt", 1941353, day1Part1},
		{"Day 1: Part 2 - Example", "example.txt", 31, day1Part2},
		{"Day 1: Part 2 - Input", "input.txt", 22539317, day1Part2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

func day1Part1(path string) int {
	file := utils.ReadFile(path)

	leftList := []int{}
	rightList := []int{}

	for _, line := range file {
		parts := strings.Fields(line)

		n1 := utils.ParseInt(parts[0])
		n2 := utils.ParseInt(parts[1])

		leftList = append(leftList, n1)
		rightList = append(rightList, n2)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	var result int
	for i, int := range leftList {
		result += utils.Abs(int - rightList[i])
	}

	return result
}

func day1Part2(path string) int {
	file := utils.ReadFile(path)

	leftList := []int{}
	numToCountsMap := map[int]int{}

	for _, line := range file {
		parts := strings.Fields(line)

		num1 := utils.ParseInt(parts[0])
		num2 := utils.ParseInt(parts[1])

		leftList = append(leftList, num1)

		numToCountsMap[num2]++
	}

	var result int
	for _, int := range leftList {
		if value, ok := numToCountsMap[int]; ok {
			result += value * int
		}
	}

	return result
}
