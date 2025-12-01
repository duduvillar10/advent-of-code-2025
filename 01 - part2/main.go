package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {

	values, err := os.ReadFile("01/input.txt")

	if err != nil {
		panic(err)
	}

	moves := []string{}
	if len(values) > 0 {
		moves = strings.Split(strings.TrimSpace(string(values)), "\n")
	}

	countPasses := 0
	const startPosition = 50
	position := startPosition

	for _, move := range moves {
		direction := move[0]
		steps, _ := strconv.Atoi(move[1:])

		if fullRotations := steps / 100; fullRotations >= 1 {
			countPasses += fullRotations
		}
		step := steps % 100

		if direction == 'L' {
			if step >= position && position != 0 {
				countPasses++
			}
			step = -step
		} else {
			if position+step >= 100 {
				countPasses++
			}
		}

		if steps == 100 {
			countPasses++
		}

		position = (position + step + 100) % 100
	}
	println(countPasses)
}
