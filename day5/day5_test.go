package day5_test

// https://adventofcode.com/2024/day/5

import (
	"sort"
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
		{"Day 5: Part 2 - Example", "example.txt", 123, day5Part2},
		{"Day 5: Part 2 - Input", "input.txt", 5273, day5Part2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

func day5Part1(path string) int {
	return calculateMiddleSum(func() ([][]int, [][]int) { return parseInput(path) }, didFollowRules, false)
}

func day5Part2(path string) int {
	return calculateMiddleSum(func() ([][]int, [][]int) { return parseInput(path) }, didFollowRules, true)
}
func didFollowRules(update []int, rules [][]int) bool {
	numToIndexMap := make(map[int]int, len(update))
	for i, num := range update {
		numToIndexMap[num] = i
	}

	for _, r := range rules {
		idx1, ok1 := numToIndexMap[r[0]]
		idx2, ok2 := numToIndexMap[r[1]]

		if (ok1 && ok2) && (idx1 >= idx2) {
			return false
		}
	}
	return true
}

func calculateMiddleSum(
	inputData func() ([][]int, [][]int),
	validate func([]int, [][]int) bool,
	reorder bool,
) int {
	rules, updates := inputData()
	sum := 0
	for _, update := range updates {
		if validate(update, rules) {
			if !reorder {
				sum += update[len(update)/2]
			}
		} else if reorder {
			corrected := reorderUpdate(update, rules)
			sum += corrected[len(corrected)/2]
		}
	}
	return sum
}

func reorderUpdate(update []int, rules [][]int) []int {
	ruleMap := make(map[int][]int)
	for _, r := range rules {
		ruleMap[r[0]] = append(ruleMap[r[0]], r[1])
	}

	sorted := append([]int(nil), update...)
	sort.SliceStable(sorted, func(i, j int) bool {
		a, b := sorted[i], sorted[j]
		for _, dependent := range ruleMap[a] {
			if dependent == b {
				return true
			}
		}

		return false
	})

	return sorted
}

func parseInput(path string) ([][]int, [][]int) {
	var rules, updates [][]int
	emptyLineFound := false

	for _, line := range utils.ReadFile(path) {
		if line == "" {
			emptyLineFound = true
			continue
		}

		if emptyLineFound {
			updates = append(updates, parseOrders(line))
		} else {
			rules = append(rules, parseRules(line))
		}
	}

	return rules, updates
}

func parseRules(line string) []int {
	parts := strings.Split(line, "|")
	return []int{utils.ParseInt(parts[0]), utils.ParseInt(parts[1])}
}

func parseOrders(line string) []int {
	parts := strings.Split(line, ",")
	order := make([]int, len(parts))

	for i, part := range parts {
		order[i] = utils.ParseInt(part)
	}

	return order
}
