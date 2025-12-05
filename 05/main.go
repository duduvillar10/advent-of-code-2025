package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	values, err := os.ReadFile("05/input.txt")
	if err != nil {
		panic(err)
	}

	groupDatabases := []string{}
	if len(values) > 0 {
		groupDatabases = strings.Split(strings.TrimSpace(string(values)), "\n\n")
	}

	rangeIds := strings.Split(strings.TrimSpace(groupDatabases[0]), "\n")
	ids := strings.Split(strings.TrimSpace(groupDatabases[1]), "\n")

	freshIDItens := [][]int{}

	for _, entry := range rangeIds {
		parts := strings.Split(entry, "-")
		start := convertToInt(parts[0])
		end := convertToInt(parts[1])

		freshIDItens = append(freshIDItens, []int{start, end})
	}

	total := 0

outer:
	for i := range ids {
		singleID := convertToInt(ids[i])

		for _, idItem := range freshIDItens {
			if singleID >= idItem[0] && singleID <= idItem[1] {
				total++
				continue outer
			}
		}
	}

	// Process ranges for p2

	newDatabase := freshIDItens

	sort.Slice(newDatabase, func(i, j int) bool {
		return newDatabase[i][0] < newDatabase[j][0]
	})

	totalp2 := 0

	for i := range newDatabase {
		current := newDatabase[i]
		totalp2 += current[1] - current[0] + 1

		if i == len(newDatabase)-1 {
			break
		}

		next := newDatabase[i+1]

		if current[1] >= next[1] {
			next[0] = current[1] + 1
			next[1] = current[1]
			continue
		}

		if current[1] >= next[0] {
			next[0] = current[1] + 1
		}
	}

	println("p1:", total)
	println("p2:", totalp2)
}

func convertToInt(s string) int {
	value, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return value
}
