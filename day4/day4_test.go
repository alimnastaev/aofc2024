package day4_test

// https://adventofcode.com/2024/day/4

import (
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
		{"Day 4: Part 1 - Example", "example.txt", 18, day4Part1},
		{"Day 4: Part 1 - Input", "input.txt", 2427, day4Part1},
		{"Day 4: Part 2 - Example", "example.txt", 9, day4Part2},
		{"Day 4: Part 2 - Input", "input.txt", 1900, day4Part2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.testFunc(tt.path))
		})
	}
}

func day4Part1(path string) int {
	lines := utils.ReadFile(path)

	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	return find1(grid)
}
func find1(grid [][]rune) int {
	word := "XMAS"
	dirs := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	rows := len(grid)
	cols := len(grid[0])
	wordLen := len(word)
	count := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range dirs {
				dr, dc := dir[0], dir[1]
				match := true

				for k := 0; k < wordLen; k++ {
					nr, nc := r+k*dr, c+k*dc
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != rune(word[k]) {
						match = false
						break
					}
				}

				if match {
					count++
				}
			}
		}
	}

	return count
}

func day4Part2(path string) int {
	lines := utils.ReadFile(path)

	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	return find2(grid)
}

func find2(grid [][]rune) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] == 'A' {
				if (grid[r-1][c-1] == 'M' && grid[r+1][c+1] == 'S') ||
					(grid[r-1][c-1] == 'S' && grid[r+1][c+1] == 'M') {
					if (grid[r-1][c+1] == 'M' && grid[r+1][c-1] == 'S') ||
						(grid[r-1][c+1] == 'S' && grid[r+1][c-1] == 'M') {
						count++
					}
				}
			}
		}
	}

	return count
}
