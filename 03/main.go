package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	values, err := os.ReadFile("03/input.txt")

	if err != nil {
		panic(err)
	}

	batteries := []string{}
	if len(values) > 0 {
		batteries = strings.Split(strings.TrimSpace(string(values)), "\n")
	}

	total := 0

	for _, battery := range batteries {
		firstLargeBattery := 0
		secondLargeBattery := 0

		for _, char := range battery {
			value := convertToInt(string(char))

			if value > secondLargeBattery {
				if secondLargeBattery > firstLargeBattery {
					firstLargeBattery = secondLargeBattery
				}
				secondLargeBattery = value
			} else if secondLargeBattery > firstLargeBattery {
				firstLargeBattery = secondLargeBattery
				secondLargeBattery = value
			}
		}

		value, err := strconv.Atoi(fmt.Sprintf("%d%d", firstLargeBattery, secondLargeBattery))
		if err != nil {
			panic(err)
		}

		total += value
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
