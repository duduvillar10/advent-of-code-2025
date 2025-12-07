package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()

	defer func() {
		println("Execution time:", time.Since(now).String())
	}()

	values, err := os.ReadFile("06/input.txt")
	if err != nil {
		panic(err)
	}

	numbers := []string{}
	if len(values) > 0 {
		numbers = strings.Split(string(values), "\n")
	}

	allDigits := [][]string{}
	for _, line := range numbers {
		allDigits = append(allDigits, strings.Split(line, ""))
	}

	operation := ""
	preSum := 0
	preProduct := 1

	total := 0

	for colIdx := 0; colIdx < len(allDigits[0]); colIdx++ {
		digits := ""

		for rowIdx := 0; rowIdx < len(allDigits)-1; rowIdx++ {
			digits += allDigits[rowIdx][colIdx]
		}

		if allDigits[len(allDigits)-1][colIdx] != " " {
			operation = allDigits[len(allDigits)-1][colIdx]
		}

		if operation == "+" {
			preSum += convertToInt(digits)
		}

		if operation == "*" {
			value := convertToInt(digits)
			if value == 0 {
				value = 1
			}
			preProduct *= value
		}

		if strings.TrimSpace(digits) == "" {
			if operation == "+" {
				total += preSum
				preSum = 0
				continue
			}

			if operation == "*" {
				total += preProduct
				preProduct = 1
			}
		}
	}

	if operation == "+" {
		total += preSum
	}

	if operation == "*" {
		total += preProduct
	}

	println("Total:", total)

}

func convertToInt(s string) int {
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return value
}
