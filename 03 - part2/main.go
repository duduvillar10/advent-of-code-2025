package main

import (
	"os"
	"strconv"
	"strings"
)

const batteryCapacity = 12

func main() {
	values, err := os.ReadFile("03 - part2/input.txt")

	if err != nil {
		panic(err)
	}

	banks := []string{}
	if len(values) > 0 {
		banks = strings.Split(strings.TrimSpace(string(values)), "\n")
	}

	total := 0

	for _, bank := range banks {
		pivotBank := bank[0:batteryCapacity]

		for i := batteryCapacity; i < len(bank); i++ {
			value := convertToInt(string(bank[i]))
			lastFromPivot := convertToInt(string(pivotBank[batteryCapacity-1]))

			newBank := removeLowestBatteryFrontToBack(pivotBank)
			if newBank != pivotBank {
				pivotBank = newBank + string(bank[i])
				continue
			}

			if value > lastFromPivot {
				pivotBank = removeLowestBatteryBackToFront(pivotBank) + string(bank[i])
			}
		}

		value, err := strconv.Atoi(pivotBank)
		if err != nil {
			panic(err)
		}

		total += value
	}

	println("p1:", total)
}

func removeLowestBatteryBackToFront(bank string) string {
	for i := len(bank) - 1; i > 0; i-- {
		current := convertToInt(string(bank[i]))
		previous := convertToInt(string(bank[i-1]))

		if current < previous {
			return bank[0:i] + bank[i+1:]
		}
	}

	return bank[0 : len(bank)-1]
}

func removeLowestBatteryFrontToBack(bank string) string {
	for i := 0; i < len(bank)-1; i++ {
		current := convertToInt(string(bank[i]))
		next := convertToInt(string(bank[i+1]))

		if current < next {
			return bank[0:i] + bank[i+1:]
		}
	}

	return bank
}

func convertToInt(s string) int {
	value, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return value
}
