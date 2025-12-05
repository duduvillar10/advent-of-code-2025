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

	database := []string{}
	if len(values) > 0 {
		database = strings.Split(strings.TrimSpace(string(values)), "\n")
	}

	freshIDItens := [][]int{}

	splitIndex := -1

	for i, entry := range database {
		if strings.Contains(entry, "-") {
			parts := strings.Split(entry, "-")
			start := convertToInt(parts[0])
			end := convertToInt(parts[1])

			freshIDItens = append(freshIDItens, []int{start, end})
			continue
		}

		if entry == "" {
			splitIndex = i
			break
		}
	}

	total := 0

outer:
	for i := splitIndex + 1; i < len(database); i++ {
		singleID := convertToInt(database[i])

		for _, idItem := range freshIDItens {
			if singleID >= idItem[0] && singleID <= idItem[1] {
				total++
				continue outer
			}
		}
	}

	// Process ranges for p2

	newDatabase := [][]int{}

	for i := 0; i < splitIndex; i++ {
		values := strings.Split(database[i], "-")
		ai := convertToInt(values[0])
		bi := convertToInt(values[1])

		newDatabase = append(newDatabase, []int{ai, bi})
	}

	sort.Slice(newDatabase, func(i, j int) bool {
		return newDatabase[i][0] < newDatabase[j][0]
	})

	totalp2 := 0

	for i := 0; i < len(newDatabase); i++ {
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
