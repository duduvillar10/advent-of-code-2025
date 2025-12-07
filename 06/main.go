package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	values, err := os.ReadFile("06/input.txt")
	if err != nil {
		panic(err)
	}

	numbers := []string{}
	if len(values) > 0 {
		numbers = strings.Split(strings.TrimSpace(string(values)), "\n")
	}

	allNumbers := [][]string{}
	for _, line := range numbers {
		parts := strings.Fields(line)
		allNumbers = append(allNumbers, parts)
	}

	header := allNumbers[len(allNumbers)-1]

	total := 0

	for colIdx := 0; colIdx < len(header); colIdx++ {
		totalColMul := 1
		totalColSum := 0

		if header[colIdx] == "+" {
			for rowIdx := 0; rowIdx < len(allNumbers)-1; rowIdx++ {
				totalColSum += convertToInt(allNumbers[rowIdx][colIdx])
			}
			total += totalColSum
		}

		if header[colIdx] == "*" {
			for rowIdx := 0; rowIdx < len(allNumbers)-1; rowIdx++ {
				totalColMul *= convertToInt(allNumbers[rowIdx][colIdx])
			}
			total += totalColMul
		}
	}

	println("p1:", total)
}

func convertToInt(s string) int {
	value, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return value
}
