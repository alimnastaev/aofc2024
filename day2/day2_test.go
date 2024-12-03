package day2_test

// https://adventofcode.com/2024/day/2

import (
	"strings"
	"sync"
	"testing"

	"github.com/alimnastaev/aofc2023/utils"
	"github.com/stretchr/testify/require"
)

// To run: go test -bench=.
func Benchmark_Day2Part1Parallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day2Part1Parallel("input.txt")
	}
}

// To run: go test -bench=.
func Benchmark_Day2Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day2Part1("input.txt")
	}
}

func Test_day2(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		want     int
		testFunc func(string) int
	}{
		{"Day 2: Part 1 - Example", "example.txt", 2, day2Part1},
		{"Day 2: Part 1 - Input", "input.txt", 663, day2Part1},
		{"Day 2: Part 2 - Example", "example.txt", 4, day2Part2},
		{"Day 2: Part 2 - Input", "input.txt", 692, day2Part2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

func day2Part1(path string) int {
	file := utils.ReadFile(path)

	var safe int
	for _, line := range file {
		if isSafe(strings.Fields(line)) {
			safe++
		}
	}

	return safe
}

func day2Part2(path string) int {
	file := utils.ReadFile(path)

	var safe int
	for _, line := range file {
		levels := strings.Fields(line)
		if isSafe(levels) {
			safe++
			continue
		}

		for i := range levels {
			clone := make([]string, len(levels))
			copy(clone, levels)

			if isSafe(append(clone[:i], clone[i+1:]...)) {
				safe++
				break
			}
		}
	}

	return safe
}

func isSafe(levels []string) bool {
	var increasing, decreasing bool
	for i := 0; i < len(levels)-1; i++ {
		n1 := utils.ParseInt(levels[i])
		n2 := utils.ParseInt(levels[i+1])

		diff := utils.Abs(n1 - n2)
		if diff < 1 || diff > 3 {
			return false
		}

		if n1 < n2 {
			if decreasing {
				return false
			}

			increasing = true
		} else if n1 > n2 {
			if increasing {
				return false
			}

			decreasing = true
		}
	}

	return true
}

func day2Part1Parallel(path string) int {
	file := utils.ReadFile(path)

	var wg sync.WaitGroup
	rowsCh := make(chan bool, len(file))

	worker := func(row string) {
		defer wg.Done()
		rowsCh <- isSafe(strings.Fields(row))
	}

	for _, line := range file {
		wg.Add(1)
		go worker(line)
	}

	go func() {
		wg.Wait()
		close(rowsCh)
	}()

	var safe int
	for result := range rowsCh {
		if result {
			safe++
		}
	}

	return safe
}
