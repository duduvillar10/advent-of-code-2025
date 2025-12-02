package main

import (
	"os"
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

			if isInvalidID(i) {
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
			if compareStrings(numberStr, len(numberStr)/idx) {
				return true
			}
		}
	}
	return false
}

func compareStrings(s string, size int) bool {
	value := s[0:size]
	for i := size; i < len(s); i += size {
		end := min(i+size, len(s))
		if s[i:end] != value {
			return false
		}
		value = s[i:end]
	}

	return true
}
