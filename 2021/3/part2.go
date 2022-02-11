package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputs []string

	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		inputs = append(inputs, input)
	}

	oxygenGeneratorRating := findOxygenGeneratorRating(inputs)
	co2ScrubberRating := findCO2ScrubberRating(inputs)

	fmt.Printf("Result: %d", oxygenGeneratorRating*co2ScrubberRating)
}

func getNumberOfZerosAndOnes(inputs []string, n int) (int, int) {
	numberOfZeros := 0
	numberOfOnes := 0

	for i, _ := range inputs {
		if inputs[i][n] == '0' {
			numberOfZeros++
		} else if inputs[i][n] == '1' {
			numberOfOnes++
		}
	}

	return numberOfZeros, numberOfOnes
}

func findOxygenGeneratorRating(inputs []string) int64 {
	remaining := inputs
	column := 0

	for len(remaining) > 1 {
		zeros, ones := getNumberOfZerosAndOnes(remaining, column)

		var newRemaining []string

		for _, number := range remaining {
			if zeros > ones && number[column] == '0' {
				newRemaining = append(newRemaining, number)
			}

			if zeros < ones && number[column] == '1' {
				newRemaining = append(newRemaining, number)
			}

			if zeros == ones && number[column] == '1' {
				newRemaining = append(newRemaining, number)
			}
		}

		remaining = newRemaining
		column++
	}

	result, _ := strconv.ParseInt(remaining[0], 2, 32)
	return result
}

func findCO2ScrubberRating(inputs []string) int64 {
	remaining := inputs
	column := 0

	for len(remaining) > 1 {
		zeros, ones := getNumberOfZerosAndOnes(remaining, column)

		var newRemaining []string

		for _, number := range remaining {
			if zeros > ones && number[column] == '1' {
				newRemaining = append(newRemaining, number)
			}

			if zeros < ones && number[column] == '0' {
				newRemaining = append(newRemaining, number)
			}

			if zeros == ones && number[column] == '0' {
				newRemaining = append(newRemaining, number)
			}
		}

		remaining = newRemaining
		column++
	}

	result, _ := strconv.ParseInt(remaining[0], 2, 32)
	return result
}
