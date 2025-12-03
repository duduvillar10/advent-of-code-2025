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

	groupBatteries := []string{}
	if len(values) > 0 {
		groupBatteries = strings.Split(strings.TrimSpace(string(values)), "\n")
	}

	total := 0

	for _, batteries := range groupBatteries {
		pivotBattery := batteries[0:batteryCapacity]

		for i := batteryCapacity; i < len(batteries); i++ {
			value := convertToInt(string(batteries[i]))
			lastPivot := convertToInt(string(pivotBattery[batteryCapacity-1]))

			newBatteries := removeLowestBatteryFrontToBack(pivotBattery)
			if newBatteries != pivotBattery {
				pivotBattery = newBatteries + string(batteries[i])
				continue
			}

			if value > lastPivot {
				pivotBattery = removeLowestBatteryBackToFront(pivotBattery) + string(batteries[i])
			}
		}

		value, err := strconv.Atoi(pivotBattery)
		if err != nil {
			panic(err)
		}

		total += value
	}

	println("p1:", total)
}

func removeLowestBatteryBackToFront(batteries string) string {
	for i := len(batteries) - 1; i > 0; i-- {
		current := convertToInt(string(batteries[i]))
		previous := convertToInt(string(batteries[i-1]))

		if current < previous {
			return batteries[0:i] + batteries[i+1:]
		}
	}

	return batteries[0 : len(batteries)-1]
}

func removeLowestBatteryFrontToBack(batteries string) string {
	for i := 0; i < len(batteries)-1; i++ {
		current := convertToInt(string(batteries[i]))
		next := convertToInt(string(batteries[i+1]))

		if current < next {
			return batteries[0:i] + batteries[i+1:]
		}
	}

	return batteries
}

func convertToInt(s string) int {
	value, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return value
}
