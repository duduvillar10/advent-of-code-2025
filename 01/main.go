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

		if direction == 'L' {
			steps = -steps
		}

		position = (position + steps + 100) % 100

		if position == 0 {
			countPasses++
		}
	}
	println(countPasses)
}
