package day1_test

// https://adventofcode.com/2024/day/1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/alimnastaev/aofc2023/utils"
	"github.com/stretchr/testify/require"
)

func Test_Day1(t *testing.T) {
	tests := []struct {
		name     string
		part     string
		path     string
		want     int
		testFunc func(string) int
	}{
		{
			name:     "Day 1: Part 1 example",
			part:     "part1",
			path:     "example1.txt",
			want:     11,
			testFunc: day1Part1,
		},
		{
			name:     "Day 1: Part 1 input",
			part:     "part1",
			path:     "input.txt",
			want:     1941353,
			testFunc: day1Part1,
		},
		{
			name:     "Day 1: Part 2 example",
			part:     "part2",
			path:     "example1.txt",
			want:     31,
			testFunc: day1Part2,
		},
		{
			name:     "Day 1: Part 2 input",
			part:     "part2",
			path:     "input.txt",
			want:     22539317,
			testFunc: day1Part2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

func day1Part1(path string) int {
	file, err := utils.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return -1
	}

	leftList := []int{}
	rightList := []int{}

	for _, line := range file {
		parts := strings.Fields(line)

		n1, err1 := strconv.Atoi(parts[0])
		n2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to integers:", err1, err2)
			return -1
		}

		leftList = append(leftList, n1)
		rightList = append(rightList, n2)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	var result int
	for i, int := range leftList {
		result += utils.Abs(int - rightList[i])
	}

	return result
}

func day1Part2(path string) int {
	file, err := utils.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return -1
	}

	leftList := []int{}
	rightListMap := map[int]int{}

	for _, line := range file {
		parts := strings.Fields(line)

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting to integers:", err1, err2)
			return -1
		}

		leftList = append(leftList, num1)

		value, ok := rightListMap[num2]
		if ok {
			rightListMap[num2] = value + 1
		} else {
			rightListMap[num2] = 1
		}
	}

	var result int
	for _, int := range leftList {
		value, ok := rightListMap[int]
		if ok {
			result += value * int
		}
	}

	return result
}
