package main

import (
	"os"
	"strings"
)

func main() {
	values, err := os.ReadFile("04/input.txt")

	if err != nil {
		panic(err)
	}

	diagram := [][]string{}
	if len(values) > 0 {
		for line := range strings.SplitSeq(strings.TrimSpace(string(values)), "\n") {
			diagram = append(diagram, strings.Split(strings.TrimSpace(line), ""))
		}
	}

	totalremovedRolls := removeRolls(diagram)
	println("p1:", totalremovedRolls)
	for {
		oldremovedRolls := totalremovedRolls
		totalremovedRolls += removeRolls(diagram)
		if oldremovedRolls == totalremovedRolls {
			break
		}
	}

	println("p2:", totalremovedRolls)
}

func removeRolls(diagram [][]string) int {

	positions := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	invalid := 0
	totalrolls := 0

	for rowIdx, row := range diagram {
		for colIdx := range row {

			if diagram[rowIdx][colIdx] != "@" {
				continue
			}

			nearRolls := 0
			totalrolls += 1

			for _, pos := range positions {
				newRow := rowIdx + pos[0]
				newCol := colIdx + pos[1]

				if newRow >= 0 && newRow < len(diagram) && newCol >= 0 && newCol < len(row) {
					if diagram[newRow][newCol] == "@" {
						nearRolls += 1
					}
				}

				if nearRolls > 3 {
					invalid += 1
					break
				}
			}

			if nearRolls <= 3 {
				diagram[rowIdx][colIdx] = "."
			}

		}
	}

	return totalrolls - invalid
}
