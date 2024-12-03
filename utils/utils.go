package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

func Abs[T Number](n T) T {
	if n < 0 {
		return -n
	}

	return n
}

func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func ReadFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
