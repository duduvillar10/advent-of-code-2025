package main

import (
	"os"
	"strings"
	"time"
)

func main() {
	now := time.Now()

	defer func() {
		println("Execution time:", time.Since(now).String())
	}()

	values, err := os.ReadFile("07/input.txt")
	if err != nil {
		panic(err)
	}

	diagram := [][]string{}
	if len(values) > 0 {
		lines := strings.Split(string(values), "\n")
		for _, line := range lines {
			diagram = append(diagram, strings.Split(line, ""))
		}
	}

	totalCount := 0

	for rowIdx := 1; rowIdx < len(diagram); rowIdx++ {
		for colIdx := (len(diagram[rowIdx]) - rowIdx) / 2; colIdx < (len(diagram[rowIdx])+rowIdx)/2; colIdx++ {
			upper := diagram[rowIdx-1][colIdx]

			if upper != "S" && upper != "|" {
				continue
			}

			current := diagram[rowIdx][colIdx]

			if current == "." {
				diagram[rowIdx][colIdx] = "|"
				continue
			}

			if current == "^" {
				totalCount++
				diagram[rowIdx][colIdx-1] = "|"
				diagram[rowIdx][colIdx+1] = "|"
			}

		}
	}

	pathWays := make(map[state]uint64)
	totalCountP2 := countPaths(diagram, 1, len(diagram[0])/2, pathWays)

	println("p1:", totalCount)
	println("p2:", totalCountP2)
}

type state struct {
	rowIdx int
	colIdx int
}

func countPaths(diagram [][]string, rowIdx, colIdx int, pathWays map[state]uint64) uint64 {
	if rowIdx >= len(diagram) {
		return 1
	}

	current := diagram[rowIdx][colIdx]
	memoKey := state{rowIdx, colIdx}

	if val, exists := pathWays[memoKey]; exists {
		return val
	}

	if current == "|" {
		count := countPaths(diagram, rowIdx+1, colIdx, pathWays)
		pathWays[memoKey] = count

		return count
	}

	if current == "^" {
		count := countPaths(diagram, rowIdx+1, colIdx-1, pathWays) + countPaths(diagram, rowIdx+1, colIdx+1, pathWays)
		pathWays[memoKey] = count

		return count
	}

	return 0
}
