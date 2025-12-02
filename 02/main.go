package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	values, err := os.ReadFile("02/input.txt")

	if err != nil {
		panic(err)
	}

	ids := []string{}
	if len(values) > 0 {
		ids = strings.Split(strings.TrimSpace(string(values)), ",")
	}

	total := 0
	for _, id := range ids {
		part1, part2 := splitID(id)

		for i := part1; i <= part2; i++ {

			isInvalid := isInvalidID(i)
			if isInvalid {
				total += i
			}
		}
	}

	println("Total:", total)
}

func splitID(id string) (int, int) {
	ids := strings.SplitN(id, "-", 2)

	integer1, err := strconv.Atoi(strings.TrimSpace(ids[0]))
	if err != nil {
		panic(err)
	}
	integer2, err := strconv.Atoi(strings.TrimSpace(ids[1]))
	if err != nil {
		panic(err)
	}
	return integer1, integer2
}

func isInvalidID(id int) bool {
	numberStr := strconv.Itoa(id)

	for idx := range numberStr {
		idx += 1

		if len(numberStr)%idx == 0 && idx != 1 {
			size := len(numberStr) / idx
			groups := splitInGroups(numberStr, size)

			if compareStrings(groups) {
				return true
			}
		}
	}
	return false
}

func compareStrings(values []string) bool {
	if len(values) < 2 {
		return false
	}

	firstValue := values[0]
	for _, value := range values[1:] {
		if value != firstValue {
			return false
		}
	}
	return true
}

func splitInGroups(s string, size int) []string {
	re := regexp.MustCompile(".{1," + strconv.Itoa(size) + "}")
	groups := re.FindAllString(s, -1)
	return groups
}
