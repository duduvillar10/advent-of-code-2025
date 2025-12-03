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

			newBank := pivotBank + string(bank[i])

			for j := 0; j < len(newBank)-1; j++ {
				current := convertToInt(string(newBank[j]))
				next := convertToInt(string(newBank[j+1]))

				if current < next {
					pivotBank = newBank[0:j] + newBank[j+1:]
					break
				}
			}
		}

		total += convertToInt(pivotBank)
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
