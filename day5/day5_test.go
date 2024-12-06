package day5_test

// https://adventofcode.com/2024/day/5

import (
	"strings"
	"testing"

	"github.com/alimnastaev/aofc2023/utils"
	"github.com/stretchr/testify/require"
)

func Test_day2(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		want     int
		testFunc func(string) int
	}{
		{"Day 5: Part 1 - Example", "example.txt", 143, day5Part1},
		{"Day 5: Part 1 - Input", "input.txt", 5639, day5Part1},
		// {"Day 5: Part 2 - Example", "example.txt", 123, day5Part2},
		// {"Day 5: Part 2 - Input", "input.txt", 692, day5Part2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

func day5Part1(path string) int {
	var rules [][]int
	var updates [][]int
	emptyLineFound := false

	for _, line := range utils.ReadFile(path) {
		if line == "" {
			emptyLineFound = true
			continue
		}

		if emptyLineFound {
			updates = append(updates, parseOrders(line)...)
		} else {
			rules = append(rules, parseRules(line)...)
		}
	}

	result := 0
	for _, update := range updates {
		if didFollowRules(update, rules) {
			result += update[len(update)/2]
		}
	}

	return result
}

func didFollowRules(update []int, rules [][]int) bool {
	updatesMap := make(map[int]int, len(update))
	for i, n := range update {
		updatesMap[n] = i
	}

	for _, r := range rules {
		index1, found1 := updatesMap[r[0]]
		index2, found2 := updatesMap[r[1]]

		if found1 && found2 && index1 >= index2 {
			return false
		}
	}

	return true
}
func parseRules(line string) [][]int {
	parts := strings.Split(line, "|")

	rule := []int{
		utils.ParseInt(parts[0]),
		utils.ParseInt(parts[1]),
	}

	return [][]int{rule}
}

func parseOrders(line string) [][]int {
	parts := strings.Split(line, ",")
	order := make([]int, len(parts))

	for i, part := range parts {
		order[i] = utils.ParseInt(part)
	}

	return [][]int{order}
}

func day5Part2(path string) int {
	file := utils.ReadFile(path)
	_ = file

	return 0
}
